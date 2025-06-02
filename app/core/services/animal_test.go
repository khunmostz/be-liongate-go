package services

import (
	"context"
	"errors"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAnimalsRepository is a mock of AnimalsRepository interface
type MockAnimalsRepository struct {
	mock.Mock
}

func (m *MockAnimalsRepository) GetAnimals(ctx context.Context) ([]domain.Animals, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Animals), args.Error(1)
}

func (m *MockAnimalsRepository) CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error) {
	args := m.Called(ctx, animal)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Animals), args.Error(1)
}

func (m *MockAnimalsRepository) GetAnimalById(ctx context.Context, id string) (*domain.Animals, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Animals), args.Error(1)
}

func (m *MockAnimalsRepository) UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error) {
	args := m.Called(ctx, id, animal)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Animals), args.Error(1)
}

func (m *MockAnimalsRepository) DeleteAnimal(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGetAnimals(t *testing.T) {
	mockRepo := new(MockAnimalsRepository)
	animalService := NewAnimalService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		expectedAnimals := []domain.Animals{
			{Id: "1", Name: "Lion"},
			{Id: "2", Name: "Tiger"},
		}

		mockRepo.On("GetAnimals", ctx).Return(expectedAnimals, nil).Once()

		result, err := animalService.GetAnimals(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedAnimals, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("database error")

		mockRepo.On("GetAnimals", ctx).Return(nil, expectedErr).Once()

		result, err := animalService.GetAnimals(ctx)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateAnimal(t *testing.T) {
	mockRepo := new(MockAnimalsRepository)
	animalService := NewAnimalService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		animal := &domain.Animals{
			Id:   "1",
			Name: "Lion",
		}

		mockRepo.On("CreateAnimal", ctx, animal).Return(animal, nil).Once()

		result, err := animalService.CreateAnimal(ctx, animal)

		assert.NoError(t, err)
		assert.Equal(t, animal, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		animal := &domain.Animals{
			Id:   "1",
			Name: "Lion",
		}

		expectedErr := errors.New("database error")
		mockRepo.On("CreateAnimal", ctx, animal).Return(nil, expectedErr).Once()

		result, err := animalService.CreateAnimal(ctx, animal)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAnimalById(t *testing.T) {
	mockRepo := new(MockAnimalsRepository)
	animalService := NewAnimalService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		animalId := "1"
		expectedAnimal := &domain.Animals{
			Id:   animalId,
			Name: "Lion",
		}

		mockRepo.On("GetAnimalById", ctx, animalId).Return(expectedAnimal, nil).Once()

		result, err := animalService.GetAnimalById(ctx, animalId)

		assert.NoError(t, err)
		assert.Equal(t, expectedAnimal, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		animalId := "999"
		expectedErr := errors.New("animal not found")

		mockRepo.On("GetAnimalById", ctx, animalId).Return(nil, expectedErr).Once()

		result, err := animalService.GetAnimalById(ctx, animalId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateAnimal(t *testing.T) {
	mockRepo := new(MockAnimalsRepository)
	animalService := NewAnimalService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		animalId := "1"
		animal := &domain.Animals{
			Id:   animalId,
			Name: "Updated Lion",
		}

		mockRepo.On("UpdateAnimal", ctx, animalId, animal).Return(animal, nil).Once()

		result, err := animalService.UpdateAnimal(ctx, animalId, animal)

		assert.NoError(t, err)
		assert.Equal(t, animal, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		animalId := "999"
		animal := &domain.Animals{
			Id:   animalId,
			Name: "Updated Lion",
		}
		expectedErr := errors.New("animal not found")

		mockRepo.On("UpdateAnimal", ctx, animalId, animal).Return(nil, expectedErr).Once()

		result, err := animalService.UpdateAnimal(ctx, animalId, animal)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteAnimal(t *testing.T) {
	mockRepo := new(MockAnimalsRepository)
	animalService := NewAnimalService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		animalId := "1"

		mockRepo.On("DeleteAnimal", ctx, animalId).Return(nil).Once()

		err := animalService.DeleteAnimal(ctx, animalId)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		animalId := "999"
		expectedErr := errors.New("animal not found")

		mockRepo.On("DeleteAnimal", ctx, animalId).Return(expectedErr).Once()

		err := animalService.DeleteAnimal(ctx, animalId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
