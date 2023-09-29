package dto

import "github.com/golang-jwt/jwt/v4"

type Payload struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

type (
	LoginRequest struct {
		Email    string `json:"email,required"`
		Password string `json:"password,required"`
	}

	RegisterRequest struct {
		Email    string `json:"email,required"`
		Password string `json:"password,required"`
	}

	LogoutRequest struct {
	}

	RefreshRequest struct {
	}
)
