package service

import (
	"context"
	"scanner/cmd/api/internal/app/dto"
	"scanner/internal/entity"
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
		})
	}
}

func TestAuthService_Logout(t *testing.T) {
	token, _ := authService.createJwtToken(&entity.User{
		UUID: "13325e91-30ca-4bbf-a6d8-cbee1e5a9624",
	})

	ctx := context.WithValue(context.Background(), "token", token)
	type args struct {
		ctx context.Context
		dto *dto.LogoutRequest
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
				ctx: ctx,
				dto: &dto.LogoutRequest{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := authService.Logout(tt.args.ctx, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
