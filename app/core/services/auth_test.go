package services

import (
	"context"
	"os"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Set up environment variables for JWT service
	os.Setenv("JWT_SECRET", "test-secret-key")
	os.Setenv("JWT_ACCESS_DURATION", "15m")
	os.Setenv("JWT_REFRESH_DURATION", "168h")
	defer func() {
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("JWT_ACCESS_DURATION")
		os.Unsetenv("JWT_REFRESH_DURATION")
	}()

	mockRepo := new(MockUsersRepository)
	mockJWT, err := utils.NewJWTService()
	assert.NoError(t, err)
	authService := NewAuthService(mockRepo, mockJWT)

	ctx := context.Background()

	// Hash the password for testing
	hashedPassword, _ := utils.HashPassword("password")

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetUserByUsername", ctx, "testuser").Return(&domain.Users{Id: "1", Username: "testuser", Password: hashedPassword}, nil).Once()

		result, err := authService.Login(ctx, &domain.LoginRequest{
			Username: "testuser",
			Password: "password",
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		mockRepo.On("GetUserByUsername", ctx, "testuser").Return(&domain.Users{Id: "1", Username: "testuser", Password: hashedPassword}, nil).Once()

		result, err := authService.Login(ctx, &domain.LoginRequest{Username: "testuser", Password: "wrongpassword"})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid password", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("GetUserByUsername", ctx, "testuser").Return(nil, nil).Once()

		result, err := authService.Login(ctx, &domain.LoginRequest{Username: "testuser", Password: "password"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}
