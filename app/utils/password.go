package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultCost is the default cost for bcrypt hashing
	DefaultCost = bcrypt.DefaultCost
	// MinCost is the minimum cost for bcrypt hashing
	MinCost = bcrypt.MinCost
	// MaxCost is the maximum cost for bcrypt hashing
	MaxCost = bcrypt.MaxCost
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrPasswordTooLong = errors.New("password is too long")
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	// Check if password is too long (bcrypt has a 72-byte limit)
	if len(password) > 72 {
		return "", ErrPasswordTooLong
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// HashPasswordWithCost hashes a password using bcrypt with custom cost
func HashPasswordWithCost(password string, cost int) (string, error) {
	// Check if password is too long (bcrypt has a 72-byte limit)
	if len(password) > 72 {
		return "", ErrPasswordTooLong
	}

	// Validate cost parameter
	if cost < MinCost || cost > MaxCost {
		cost = DefaultCost
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// VerifyPassword compares a hashed password with a plain text password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// CheckPasswordStrength checks if password meets minimum requirements
func CheckPasswordStrength(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	if len(password) > 72 {
		return ErrPasswordTooLong
	}

	// You can add more password strength validations here:
	// - Must contain uppercase letter
	// - Must contain lowercase letter
	// - Must contain number
	// - Must contain special character
	// etc.

	return nil
}

// IsPasswordValid validates password and returns true if valid
func IsPasswordValid(hashedPassword, password string) bool {
	err := VerifyPassword(hashedPassword, password)
	return err == nil
}
