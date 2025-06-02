package services

import (
	"context"
	"errors"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPerformanceStageRepository is a mock of PerformanceStageRepository interface
type MockPerformanceStageRepository struct {
	mock.Mock
}

func (m *MockPerformanceStageRepository) GetStages(ctx context.Context) ([]domain.PerformanceStage, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.PerformanceStage), args.Error(1)
}

func (m *MockPerformanceStageRepository) CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	args := m.Called(ctx, stage)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PerformanceStage), args.Error(1)
}

func (m *MockPerformanceStageRepository) GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PerformanceStage), args.Error(1)
}

func (m *MockPerformanceStageRepository) UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	args := m.Called(ctx, id, stage)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PerformanceStage), args.Error(1)
}

func (m *MockPerformanceStageRepository) DeleteStage(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGetStages(t *testing.T) {
	mockRepo := new(MockPerformanceStageRepository)
	stageService := NewPerformanceStageService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		expectedStages := []domain.PerformanceStage{
			{Id: "1", RoomNumber: "A101", SeatCapacity: 100, PricePerSeat: 50.0},
			{Id: "2", RoomNumber: "B202", SeatCapacity: 200, PricePerSeat: 75.0},
		}

		mockRepo.On("GetStages", ctx).Return(expectedStages, nil).Once()

		result, err := stageService.GetStages(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedStages, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("database error")

		mockRepo.On("GetStages", ctx).Return(nil, expectedErr).Once()

		result, err := stageService.GetStages(ctx)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateStage(t *testing.T) {
	mockRepo := new(MockPerformanceStageRepository)
	stageService := NewPerformanceStageService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		stage := &domain.PerformanceStage{
			Id:           "1",
			RoomNumber:   "A101",
			SeatCapacity: 100,
			PricePerSeat: 50.0,
		}

		mockRepo.On("CreateStage", ctx, stage).Return(stage, nil).Once()

		result, err := stageService.CreateStage(ctx, stage)

		assert.NoError(t, err)
		assert.Equal(t, stage, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		stage := &domain.PerformanceStage{
			Id:           "1",
			RoomNumber:   "A101",
			SeatCapacity: 100,
			PricePerSeat: 50.0,
		}

		expectedErr := errors.New("database error")
		mockRepo.On("CreateStage", ctx, stage).Return(nil, expectedErr).Once()

		result, err := stageService.CreateStage(ctx, stage)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetStageById(t *testing.T) {
	mockRepo := new(MockPerformanceStageRepository)
	stageService := NewPerformanceStageService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		stageId := "1"
		expectedStage := &domain.PerformanceStage{
			Id:           stageId,
			RoomNumber:   "A101",
			SeatCapacity: 100,
			PricePerSeat: 50.0,
		}

		mockRepo.On("GetStageById", ctx, stageId).Return(expectedStage, nil).Once()

		result, err := stageService.GetStageById(ctx, stageId)

		assert.NoError(t, err)
		assert.Equal(t, expectedStage, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		stageId := "999"
		expectedErr := errors.New("stage not found")

		mockRepo.On("GetStageById", ctx, stageId).Return(nil, expectedErr).Once()

		result, err := stageService.GetStageById(ctx, stageId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateStage(t *testing.T) {
	mockRepo := new(MockPerformanceStageRepository)
	stageService := NewPerformanceStageService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		stageId := "1"
		stage := &domain.PerformanceStage{
			Id:           stageId,
			RoomNumber:   "A101-Updated",
			SeatCapacity: 150,
			PricePerSeat: 60.0,
		}

		mockRepo.On("UpdateStage", ctx, stageId, stage).Return(stage, nil).Once()

		result, err := stageService.UpdateStage(ctx, stageId, stage)

		assert.NoError(t, err)
		assert.Equal(t, stage, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		stageId := "999"
		stage := &domain.PerformanceStage{
			Id:           stageId,
			RoomNumber:   "A101-Updated",
			SeatCapacity: 150,
			PricePerSeat: 60.0,
		}
		expectedErr := errors.New("stage not found")

		mockRepo.On("UpdateStage", ctx, stageId, stage).Return(nil, expectedErr).Once()

		result, err := stageService.UpdateStage(ctx, stageId, stage)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteStage(t *testing.T) {
	mockRepo := new(MockPerformanceStageRepository)
	stageService := NewPerformanceStageService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		stageId := "1"

		mockRepo.On("DeleteStage", ctx, stageId).Return(nil).Once()

		err := stageService.DeleteStage(ctx, stageId)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		stageId := "999"
		expectedErr := errors.New("stage not found")

		mockRepo.On("DeleteStage", ctx, stageId).Return(expectedErr).Once()

		err := stageService.DeleteStage(ctx, stageId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
