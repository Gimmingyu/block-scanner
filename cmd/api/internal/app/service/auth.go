package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"scanner/cmd/api/internal/app/dto"
	"scanner/internal/entity"
	"scanner/pkg/repository"
	"time"
)

type AuthService struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthService(db *gorm.DB, redis *redis.Client) *AuthService {
	return &AuthService{db: db, redis: redis}
}

func (a *AuthService) createJwtToken(ctx context.Context, user *entity.User) (string, error) {
	claims := &dto.JwtPayload{
		UUID: user.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "primrose",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	signature := os.Getenv("JWT_SIGNATURE")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signature))
}

func (a *AuthService) Login(ctx context.Context, dto *dto.LoginRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		where := map[string]interface{}{"email": dto.Email}
		user, err := repository.FindOne[entity.User](tx, where)
		if err != nil {
			return fmt.Errorf("user not found: %v", err)
		}

		if err = user.ComparePassword(dto.Password); err != nil {
			return err
		}

		token, err := a.createJwtToken(ctx, user)
		if err != nil {
			return err
		}

		return a.redis.SetEx(ctx, user.UUID, token, time.Hour*24).Err()
	})
}

func (a *AuthService) Register(ctx context.Context, dto *dto.RegisterRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		where := map[string]interface{}{"email": dto.Email}
		_, err := repository.FindOne[entity.User](tx, where)
		switch {
		case err == nil:
			log.Println("HERE")
			return fmt.Errorf("user already exists: %v", err)
		case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
			return fmt.Errorf("error while finding user: %v", err)
		}
		user := &entity.User{
			Email:    dto.Email,
			Password: dto.Password,
		}
		return repository.Insert[entity.User](tx, user)
	})
}

func (a *AuthService) Logout(ctx context.Context, dto *dto.LogoutRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return nil
	})
}

func (a *AuthService) Refresh(ctx context.Context, dto *dto.RefreshRequest) error {
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
