package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"os"
	"scanner/cmd/api/internal/app/dto"
	"scanner/cmd/api/internal/app/presenter"
	"strings"
)

func Authenticated() gin.HandlerFunc {
	issuer, ok := os.LookupEnv("JWT_ISSUER")
	if !ok {
		panic("JWT_ISSUER not set")
	}

	signature, ok := os.LookupEnv("JWT_SIGNATURE")
	if !ok {
		panic("JWT_SIGNATURE not set")
	}

	return func(ctx *gin.Context) {

		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			presenter.Error(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		log.Println(authorization)
		authorization = strings.TrimPrefix(authorization, "Bearer ")
		log.Println(authorization)

		var payload dto.Payload
		token, err := jwt.ParseWithClaims(authorization, &payload, func(token *jwt.Token) (interface{}, error) {
			return []byte(signature), nil
		})

		if err != nil || !token.Valid {
			presenter.Error(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		if payload.Issuer != issuer {
			presenter.Error(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		ctx.Set("payload", payload)

		ctx.Next()
	}
}
