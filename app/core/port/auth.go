package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error)
	Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error)
	RefreshToken(ctx context.Context, req *domain.RefreshTokenRequest) (*domain.TokenPair, error)
}
