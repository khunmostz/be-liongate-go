package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrExpiredToken  = errors.New("token expired")
	ErrInvalidClaims = errors.New("invalid claims")
)

type JWTService struct {
	secretKey            []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

// NewJWTService creates a new JWT service
func NewJWTService() (*JWTService, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}

	accessDuration, err := time.ParseDuration(os.Getenv("JWT_ACCESS_DURATION")) // 15 minutes
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_ACCESS_DURATION: %w", err)
	}

	refreshDuration, err := time.ParseDuration(os.Getenv("JWT_REFRESH_DURATION")) // 7 days
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_DURATION: %w", err)
	}

	return &JWTService{
		secretKey:            []byte(secretKey),
		accessTokenDuration:  accessDuration,
		refreshTokenDuration: refreshDuration,
	}, nil
}

// GenerateAccessToken creates a new access token
func (j *JWTService) GenerateAccessToken(user *domain.Users) (string, error) {
	return j.generateToken(user, AccessTokenType, j.accessTokenDuration)
}

// GenerateRefreshToken creates a new refresh token
func (j *JWTService) GenerateRefreshToken(user *domain.Users) (string, error) {
	return j.generateToken(user, RefreshTokenType, j.refreshTokenDuration)
}

// GenerateTokenPair creates both access and refresh tokens
func (j *JWTService) GenerateTokenPair(user *domain.Users) (*domain.TokenPair, error) {
	accessToken, err := j.GenerateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := j.GenerateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// generateToken creates a JWT token with specified type and duration
func (j *JWTService) generateToken(user *domain.Users, tokenType string, duration time.Duration) (string, error) {
	now := time.Now()
	expiresAt := now.Add(duration)

	claims := domain.JWTClaims{
		UserID:    user.Id,
		Username:  user.Username,
		Role:      user.Role,
		Type:      tokenType,
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  claims.UserID,
		"username": claims.Username,
		"role":     claims.Role,
		"type":     claims.Type,
		"iat":      claims.IssuedAt,
		"exp":      claims.ExpiresAt,
	})

	return token.SignedString(j.secretKey)
}

// VerifyToken verifies and parses a JWT token
func (j *JWTService) VerifyToken(tokenString string) (*domain.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	// Check expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, ErrExpiredToken
		}
	}

	// Extract claims
	jwtClaims := &domain.JWTClaims{}

	if userID, ok := claims["user_id"].(string); ok {
		jwtClaims.UserID = userID
	}

	if username, ok := claims["username"].(string); ok {
		jwtClaims.Username = username
	}

	if role, ok := claims["role"].(string); ok {
		jwtClaims.Role = role
	}

	if tokenType, ok := claims["type"].(string); ok {
		jwtClaims.Type = tokenType
	}

	if iat, ok := claims["iat"].(float64); ok {
		jwtClaims.IssuedAt = int64(iat)
	}

	if exp, ok := claims["exp"].(float64); ok {
		jwtClaims.ExpiresAt = int64(exp)
	}

	return jwtClaims, nil
}

// VerifyAccessToken specifically verifies an access token
func (j *JWTService) VerifyAccessToken(tokenString string) (*domain.JWTClaims, error) {
	claims, err := j.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.Type != AccessTokenType {
		return nil, errors.New("not an access token")
	}

	return claims, nil
}

// VerifyRefreshToken specifically verifies a refresh token
func (j *JWTService) VerifyRefreshToken(tokenString string) (*domain.JWTClaims, error) {
	claims, err := j.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.Type != RefreshTokenType {
		return nil, errors.New("not a refresh token")
	}

	return claims, nil
}

// CreateRefreshTokenRecord creates a refresh token record
func (j *JWTService) CreateRefreshTokenRecord(userID, token string) *domain.RefreshToken {
	return &domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(j.refreshTokenDuration),
		CreatedAt: time.Now(),
	}
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
