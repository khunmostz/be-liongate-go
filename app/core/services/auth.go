package services

import (
	"context"
	"errors"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/utils"
)

type AuthService struct {
	userRepo   port.UsersRepository
	jwtService *utils.JWTService
}

func NewAuthService(userRepo port.UsersRepository, jwtService *utils.JWTService) *AuthService {
	return &AuthService{userRepo: userRepo, jwtService: jwtService}
}

func (s *AuthService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	valid := utils.IsPasswordValid(user.Password, req.Password)
	if !valid {
		return nil, errors.New("invalid password")
	}

	access_token, err := s.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refresh_token, err := s.jwtService.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User: user,
		Tokens: &domain.TokenPair{
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		},
	}, nil
}

func (s *AuthService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.Users{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if _, err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	access_token, err := s.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refresh_token, err := s.jwtService.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User: user,
		Tokens: &domain.TokenPair{
			AccessToken:  access_token,
			RefreshToken: refresh_token,
		},
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *domain.RefreshTokenRequest) (*domain.TokenPair, error) {
	claims, err := s.jwtService.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserById(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  token,
		RefreshToken: req.RefreshToken,
	}, nil
}
