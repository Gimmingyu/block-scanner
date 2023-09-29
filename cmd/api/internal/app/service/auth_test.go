package service

import (
	"context"
	"log"
	"scanner/cmd/api/internal/app/dto"
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *dto.RegisterRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success case",
			args: args{
				ctx: context.Background(),
				dto: &dto.RegisterRequest{
					Email:    "example@gmail.com",
					Password: "example12#$",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := authService.Register(tt.args.ctx, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		dto *dto.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success case",
			args: args{
				ctx: context.Background(),
				dto: &dto.LoginRequest{
					Email:    "example@gmail.com",
					Password: "example12#$",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := authService.Login(tt.args.ctx, tt.args.dto); err != nil {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			log.Println(tt.args.ctx.Value("token"))
		})
	}
}
