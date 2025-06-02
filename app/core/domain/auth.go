package domain

import "time"

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the register request payload
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// TokenPair represents access and refresh tokens
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	User   *Users     `json:"user"`
	Tokens *TokenPair `json:"tokens"`
}

// RefreshTokenRequest represents refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// JWTClaims represents the JWT claims structure
type JWTClaims struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Type      string `json:"type"` // "access" or "refresh"
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

// RefreshToken represents stored refresh token
type RefreshToken struct {
	ID        string    `json:"id" bson:"_id" gorm:"primaryKey;column:id;type:string"`
	UserID    string    `json:"user_id" bson:"user_id" gorm:"column:user_id;index"`
	Token     string    `json:"token" bson:"token" gorm:"column:token;uniqueIndex"`
	ExpiresAt time.Time `json:"expires_at" bson:"expires_at" gorm:"column:expires_at"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" gorm:"column:created_at"`
}
