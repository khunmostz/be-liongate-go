package services

import (
	"context"
	"errors"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockShowRoundsRepository is a mock of ShowRoundsRepository interface
type MockShowRoundsRepository struct {
	mock.Mock
}

func (m *MockShowRoundsRepository) CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	args := m.Called(ctx, showRound)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ShowRounds), args.Error(1)
}

func (m *MockShowRoundsRepository) GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ShowRounds), args.Error(1)
}

func (m *MockShowRoundsRepository) GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ShowRounds), args.Error(1)
}

func (m *MockShowRoundsRepository) UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	args := m.Called(ctx, id, showRound)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ShowRounds), args.Error(1)
}

func (m *MockShowRoundsRepository) DeleteShowRound(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateShowRound(t *testing.T) {
	mockRepo := new(MockShowRoundsRepository)
	showRoundService := NewShowRoundService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		showRound := &domain.ShowRounds{
			Id:       "1",
			AnimalId: "animal1",
			StageId:  "stage1",
			ShowTime: "2023-06-15T14:00:00Z",
		}

		mockRepo.On("CreateShowRound", ctx, showRound).Return(showRound, nil).Once()

		result, err := showRoundService.CreateShowRound(ctx, showRound)

		assert.NoError(t, err)
		assert.Equal(t, showRound, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		showRound := &domain.ShowRounds{
			Id:       "1",
			AnimalId: "animal1",
			StageId:  "stage1",
			ShowTime: "2023-06-15T14:00:00Z",
		}

		expectedErr := errors.New("database error")
		mockRepo.On("CreateShowRound", ctx, showRound).Return(nil, expectedErr).Once()

		result, err := showRoundService.CreateShowRound(ctx, showRound)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllShowRounds(t *testing.T) {
	mockRepo := new(MockShowRoundsRepository)
	showRoundService := NewShowRoundService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		showRounds := []*domain.ShowRounds{
			{
				Id: "1",
			},
		}
		mockRepo.On("GetAllShowRounds", ctx).Return(showRounds, nil).Once()

		result, err := showRoundService.GetAllShowRounds(ctx)

		assert.NoError(t, err)
		assert.Equal(t, showRounds, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockRepo.On("GetAllShowRounds", ctx).Return(nil, errors.New("database error")).Once()

		result, err := showRoundService.GetAllShowRounds(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetShowRoundById(t *testing.T) {
	mockRepo := new(MockShowRoundsRepository)
	showRoundService := NewShowRoundService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		roundId := "1"
		expectedShowRound := &domain.ShowRounds{
			Id:       roundId,
			AnimalId: "animal1",
			StageId:  "stage1",
			ShowTime: "2023-06-15T14:00:00Z",
		}

		mockRepo.On("GetShowRoundById", ctx, roundId).Return(expectedShowRound, nil).Once()

		result, err := showRoundService.GetShowRoundById(ctx, roundId)

		assert.NoError(t, err)
		assert.Equal(t, expectedShowRound, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		roundId := "999"
		expectedErr := errors.New("show round not found")

		mockRepo.On("GetShowRoundById", ctx, roundId).Return(nil, expectedErr).Once()

		result, err := showRoundService.GetShowRoundById(ctx, roundId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateShowRound(t *testing.T) {
	mockRepo := new(MockShowRoundsRepository)
	showRoundService := NewShowRoundService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		roundId := "1"
		showRound := &domain.ShowRounds{
			Id:       roundId,
			AnimalId: "animal1",
			StageId:  "stage2",
			ShowTime: "2023-06-15T16:00:00Z",
		}

		mockRepo.On("UpdateShowRound", ctx, roundId, showRound).Return(showRound, nil).Once()

		result, err := showRoundService.UpdateShowRound(ctx, roundId, showRound)

		assert.NoError(t, err)
		assert.Equal(t, showRound, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		roundId := "999"
		showRound := &domain.ShowRounds{
			Id:       roundId,
			AnimalId: "animal1",
			StageId:  "stage2",
			ShowTime: "2023-06-15T16:00:00Z",
		}
		expectedErr := errors.New("show round not found")

		mockRepo.On("UpdateShowRound", ctx, roundId, showRound).Return(nil, expectedErr).Once()

		result, err := showRoundService.UpdateShowRound(ctx, roundId, showRound)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteShowRound(t *testing.T) {
	mockRepo := new(MockShowRoundsRepository)
	showRoundService := NewShowRoundService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		roundId := "1"

		mockRepo.On("DeleteShowRound", ctx, roundId).Return(nil).Once()

		err := showRoundService.DeleteShowRound(ctx, roundId)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		roundId := "999"
		expectedErr := errors.New("show round not found")

		mockRepo.On("DeleteShowRound", ctx, roundId).Return(expectedErr).Once()

		err := showRoundService.DeleteShowRound(ctx, roundId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
