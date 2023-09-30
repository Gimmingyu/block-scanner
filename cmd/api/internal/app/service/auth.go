package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"os"
	"scanner/cmd/api/internal/app/dto"
	"scanner/internal/entity"
	"scanner/pkg/repository"
	"time"
)

type AuthService struct {
	db        *gorm.DB
	redis     *redis.Client
	signature string
	issuer    string
}

func NewAuthService(db *gorm.DB, redis *redis.Client) *AuthService {
	return &AuthService{
		db:        db,
		redis:     redis,
		signature: os.Getenv("JWT_SIGNATURE"),
		issuer:    os.Getenv("JWT_ISSUER"),
	}
}

func (a *AuthService) createJwtToken(user *entity.User) (string, error) {
	claims := &dto.Payload{
		UUID: user.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.signature))
}

func (a *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (token string, err error) {
	if err = a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var (
			where = map[string]interface{}{"email": req.Email}
			user  *entity.User
		)

		user, err = repository.FindOne[entity.User](tx, where)
		if err != nil {
			return fmt.Errorf("user not found: %v", err)
		}

		if err = user.ComparePassword(req.Password); err != nil {
			return err
		}

		token, err = a.createJwtToken(user)
		if err != nil {
			return err
		}

		return a.redis.SetEx(ctx, user.UUID, token, time.Hour*24).Err()
	}); err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		where := map[string]interface{}{"email": req.Email}
		_, err := repository.FindOne[entity.User](tx, where)
		switch {
		case err == nil:
			return fmt.Errorf("user already exists: %v", err)
		case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
			return fmt.Errorf("error while finding user: %v", err)
		}
		user := &entity.User{
			Email:    req.Email,
			Password: req.Password,
		}
		return repository.Insert[entity.User](tx, user)
	})
}

func (a *AuthService) Logout(ctx context.Context, req *dto.LogoutRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		payload, ok := ctx.Value("payload").(dto.Payload)
		if !ok {
			return errors.New("no JWT payload in context")
		}

		user, err := repository.FindOne[entity.User](tx, map[string]interface{}{"uuid": payload.UUID})
		if err != nil {
			return fmt.Errorf("user not found: %v", err)
		}

		return a.redis.Del(ctx, user.UUID).Err()
	})
}

func (a *AuthService) Refresh(ctx context.Context, req *dto.RefreshRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		payload, ok := ctx.Value("payload").(dto.Payload)
		if !ok {
			return errors.New("no JWT payload in context")
		}

		user, err := repository.FindOne[entity.User](tx, map[string]interface{}{"uuid": payload.UUID})
		if err != nil {
			return fmt.Errorf("user not found: %v", err)
		}

		token, err := a.createJwtToken(user)
		if err != nil {
			return err
		}

		return a.redis.SetEx(ctx, user.UUID, token, time.Hour*24).Err()
	})
}
