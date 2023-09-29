package dto

type (
	JwtPayload struct {
		UUID string
	}
)

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
