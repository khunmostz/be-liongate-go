package services

import (
	"context"
	"errors"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUsersRepository is a mock of UsersRepository interface
type MockUsersRepository struct {
	mock.Mock
}

func (m *MockUsersRepository) CreateUser(ctx context.Context, user *domain.Users) (*domain.Users, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Users), args.Error(1)
}

func (m *MockUsersRepository) GetUserById(ctx context.Context, id string) (*domain.Users, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Users), args.Error(1)
}

func (m *MockUsersRepository) GetUsersByRole(ctx context.Context, role string) ([]domain.Users, error) {
	args := m.Called(ctx, role)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Users), args.Error(1)
}

func (m *MockUsersRepository) GetUserByUsername(ctx context.Context, username string) (*domain.Users, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Users), args.Error(1)
}

func (m *MockUsersRepository) UpdateUser(ctx context.Context, id string, user *domain.Users) (*domain.Users, error) {
	args := m.Called(ctx, id, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Users), args.Error(1)
}

func (m *MockUsersRepository) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	userService := NewUsersService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		user := &domain.Users{
			Id:       "1",
			Username: "testuser",
			Password: "password",
			Role:     "user",
		}

		mockRepo.On("CreateUser", ctx, user).Return(user, nil).Once()

		result, err := userService.Register(ctx, user)

		assert.NoError(t, err)
		assert.Equal(t, user, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		user := &domain.Users{
			Id:       "1",
			Username: "testuser",
			Password: "password",
			Role:     "user",
		}

		expectedErr := errors.New("database error")
		mockRepo.On("CreateUser", ctx, user).Return(nil, expectedErr).Once()

		result, err := userService.Register(ctx, user)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	userService := NewUsersService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userId := "1"
		expectedUser := &domain.Users{
			Id:       userId,
			Username: "testuser",
			Password: "password",
			Role:     "user",
		}

		mockRepo.On("GetUserById", ctx, userId).Return(expectedUser, nil).Once()

		result, err := userService.GetUserById(ctx, userId)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		userId := "999"
		expectedErr := errors.New("user not found")

		mockRepo.On("GetUserById", ctx, userId).Return(nil, expectedErr).Once()

		result, err := userService.GetUserById(ctx, userId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUsersByRole(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	userService := NewUsersService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		role := "admin"
		expectedUsers := []domain.Users{
			{Id: "1", Username: "admin1", Role: "admin"},
			{Id: "2", Username: "admin2", Role: "admin"},
		}

		mockRepo.On("GetUsersByRole", ctx, role).Return(expectedUsers, nil).Once()

		result, err := userService.GetUsersByRole(ctx, role)

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		role := "unknown"
		expectedErr := errors.New("database error")

		mockRepo.On("GetUsersByRole", ctx, role).Return(nil, expectedErr).Once()

		result, err := userService.GetUsersByRole(ctx, role)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	userService := NewUsersService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userId := "1"
		user := &domain.Users{
			Id:       userId,
			Username: "updated",
			Password: "newpassword",
			Role:     "admin",
		}

		mockRepo.On("UpdateUser", ctx, userId, user).Return(user, nil).Once()

		result, err := userService.UpdateUser(ctx, userId, user)

		assert.NoError(t, err)
		assert.Equal(t, user, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userId := "999"
		user := &domain.Users{
			Id:       userId,
			Username: "updated",
			Password: "newpassword",
			Role:     "admin",
		}
		expectedErr := errors.New("user not found")

		mockRepo.On("UpdateUser", ctx, userId, user).Return(nil, expectedErr).Once()

		result, err := userService.UpdateUser(ctx, userId, user)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	userService := NewUsersService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userId := "1"

		mockRepo.On("DeleteUser", ctx, userId).Return(nil).Once()

		err := userService.DeleteUser(ctx, userId)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userId := "999"
		expectedErr := errors.New("user not found")

		mockRepo.On("DeleteUser", ctx, userId).Return(expectedErr).Once()

		err := userService.DeleteUser(ctx, userId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
