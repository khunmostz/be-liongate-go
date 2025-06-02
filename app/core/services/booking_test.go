package services

import (
	"context"
	"errors"
	"testing"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookingsRepository is a mock of BookingsRepository interface
type MockBookingsRepository struct {
	mock.Mock
}

func (m *MockBookingsRepository) CreateBooking(ctx context.Context, booking *domain.Bookings) (*domain.Bookings, error) {
	args := m.Called(ctx, booking)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bookings), args.Error(1)
}

func (m *MockBookingsRepository) GetBookingById(ctx context.Context, id string) (*domain.Bookings, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bookings), args.Error(1)
}

func (m *MockBookingsRepository) GetBookingsByUserId(ctx context.Context, userId string) ([]domain.Bookings, error) {
	args := m.Called(ctx, userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Bookings), args.Error(1)
}

func (m *MockBookingsRepository) GetBookingsByRoundId(ctx context.Context, roundId string) ([]domain.Bookings, error) {
	args := m.Called(ctx, roundId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Bookings), args.Error(1)
}

func (m *MockBookingsRepository) UpdateBooking(ctx context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error) {
	args := m.Called(ctx, id, booking)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Bookings), args.Error(1)
}

func (m *MockBookingsRepository) DeleteBooking(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateBooking(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		booking := &domain.Bookings{
			Id:         "1",
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		// Mock GetBookingsByRoundId to return empty list (no existing bookings)
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return([]domain.Bookings{}, nil).Once()

		mockRepo.On("CreateBooking", ctx, booking).Return(booking, nil).Once()

		result, err := bookingService.CreateBooking(ctx, booking)

		assert.NoError(t, err)
		assert.Equal(t, booking, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("duplicate seat number", func(t *testing.T) {
		booking := &domain.Bookings{
			Id:         "2",
			UserId:     "user2",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		// Mock GetBookingsByRoundId to return a booking with the same seat number
		existingBookings := []domain.Bookings{
			{
				Id:         "1",
				UserId:     "user1",
				RoundId:    "round1",
				SeatNumber: 5,
			},
		}
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return(existingBookings, nil).Once()

		result, err := bookingService.CreateBooking(ctx, booking)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "seat number 5 is already taken")
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error getting bookings", func(t *testing.T) {
		booking := &domain.Bookings{
			Id:         "2",
			UserId:     "user2",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		expectedErr := errors.New("database error")
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return(nil, expectedErr).Once()

		result, err := bookingService.CreateBooking(ctx, booking)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error creating booking", func(t *testing.T) {
		booking := &domain.Bookings{
			Id:         "1",
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		// Mock GetBookingsByRoundId to return empty list (no existing bookings)
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return([]domain.Bookings{}, nil).Once()

		expectedErr := errors.New("database error")
		mockRepo.On("CreateBooking", ctx, booking).Return(nil, expectedErr).Once()

		result, err := bookingService.CreateBooking(ctx, booking)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetBookingById(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		bookingId := "1"
		expectedBooking := &domain.Bookings{
			Id:     bookingId,
			UserId: "user1",
		}

		mockRepo.On("GetBookingById", ctx, bookingId).Return(expectedBooking, nil).Once()

		result, err := bookingService.GetBookingById(ctx, bookingId)

		assert.NoError(t, err)
		assert.Equal(t, expectedBooking, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		bookingId := "999"
		expectedErr := errors.New("booking not found")

		mockRepo.On("GetBookingById", ctx, bookingId).Return(nil, expectedErr).Once()

		result, err := bookingService.GetBookingById(ctx, bookingId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetBookingsByUserId(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		userId := "user1"
		expectedBookings := []domain.Bookings{
			{Id: "1", UserId: userId},
			{Id: "2", UserId: userId},
		}

		mockRepo.On("GetBookingsByUserId", ctx, userId).Return(expectedBookings, nil).Once()

		result, err := bookingService.GetBookingsByUserId(ctx, userId)

		assert.NoError(t, err)
		assert.Equal(t, expectedBookings, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userId := "user1"
		expectedErr := errors.New("database error")

		mockRepo.On("GetBookingsByUserId", ctx, userId).Return(nil, expectedErr).Once()

		result, err := bookingService.GetBookingsByUserId(ctx, userId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetBookingsByRoundId(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		roundId := "round1"
		expectedBookings := []domain.Bookings{
			{Id: "1", RoundId: roundId, SeatNumber: 1},
			{Id: "2", RoundId: roundId, SeatNumber: 2},
		}

		mockRepo.On("GetBookingsByRoundId", ctx, roundId).Return(expectedBookings, nil).Once()

		result, err := bookingService.GetBookingsByRoundId(ctx, roundId)

		assert.NoError(t, err)
		assert.Equal(t, expectedBookings, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		roundId := "round1"
		expectedErr := errors.New("database error")

		mockRepo.On("GetBookingsByRoundId", ctx, roundId).Return(nil, expectedErr).Once()

		result, err := bookingService.GetBookingsByRoundId(ctx, roundId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateBooking(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		bookingId := "1"
		booking := &domain.Bookings{
			Id:         bookingId,
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 10,
		}

		// Mock GetBookingsByRoundId to return empty list (no other bookings)
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return([]domain.Bookings{}, nil).Once()

		mockRepo.On("UpdateBooking", ctx, bookingId, booking).Return(booking, nil).Once()

		result, err := bookingService.UpdateBooking(ctx, bookingId, booking)

		assert.NoError(t, err)
		assert.Equal(t, booking, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("duplicate seat number", func(t *testing.T) {
		bookingId := "1"
		booking := &domain.Bookings{
			Id:         bookingId,
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		// Mock GetBookingsByRoundId to return a booking with the same seat number but different ID
		existingBookings := []domain.Bookings{
			{
				Id:         "2", // Different ID
				UserId:     "user2",
				RoundId:    "round1",
				SeatNumber: 5, // Same seat number
			},
		}
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return(existingBookings, nil).Once()

		result, err := bookingService.UpdateBooking(ctx, bookingId, booking)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "seat number 5 is already taken")
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("same booking different seat", func(t *testing.T) {
		bookingId := "1"
		booking := &domain.Bookings{
			Id:         bookingId,
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 10, // New seat number
		}

		// Mock GetBookingsByRoundId to return the same booking with different seat
		existingBookings := []domain.Bookings{
			{
				Id:         bookingId, // Same ID
				UserId:     "user1",
				RoundId:    "round1",
				SeatNumber: 5, // Old seat number
			},
		}
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return(existingBookings, nil).Once()
		mockRepo.On("UpdateBooking", ctx, bookingId, booking).Return(booking, nil).Once()

		result, err := bookingService.UpdateBooking(ctx, bookingId, booking)

		assert.NoError(t, err)
		assert.Equal(t, booking, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error getting bookings", func(t *testing.T) {
		bookingId := "1"
		booking := &domain.Bookings{
			Id:         bookingId,
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		expectedErr := errors.New("database error")
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return(nil, expectedErr).Once()

		result, err := bookingService.UpdateBooking(ctx, bookingId, booking)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error updating booking", func(t *testing.T) {
		bookingId := "1"
		booking := &domain.Bookings{
			Id:         bookingId,
			UserId:     "user1",
			RoundId:    "round1",
			SeatNumber: 5,
		}

		// Mock GetBookingsByRoundId to return empty list (no other bookings)
		mockRepo.On("GetBookingsByRoundId", ctx, "round1").Return([]domain.Bookings{}, nil).Once()

		expectedErr := errors.New("database error")
		mockRepo.On("UpdateBooking", ctx, bookingId, booking).Return(nil, expectedErr).Once()

		result, err := bookingService.UpdateBooking(ctx, bookingId, booking)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteBooking(t *testing.T) {
	mockRepo := new(MockBookingsRepository)
	bookingService := NewBookingsService(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		bookingId := "1"

		mockRepo.On("DeleteBooking", ctx, bookingId).Return(nil).Once()

		err := bookingService.DeleteBooking(ctx, bookingId)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		bookingId := "999"
		expectedErr := errors.New("booking not found")

		mockRepo.On("DeleteBooking", ctx, bookingId).Return(expectedErr).Once()

		err := bookingService.DeleteBooking(ctx, bookingId)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
