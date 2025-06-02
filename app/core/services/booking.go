package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type BookingService struct {
	bookingsRepository port.BookingsRepository
}

func NewBookingsService(bookingsRepository port.BookingsRepository) *BookingService {
	return &BookingService{
		bookingsRepository: bookingsRepository,
	}
}

// checkSeatAvailability checks if the seat number is already taken for a specific round
func (s *BookingService) checkSeatAvailability(ctx context.Context, roundId string, seatNumber int, excludeBookingId string) error {
	// Get all bookings for this round
	bookings, err := s.GetBookingsByRoundId(ctx, roundId)
	if err != nil {
		return err
	}

	// Check if the seat number is already taken
	for _, booking := range bookings {
		// Skip the current booking if we're updating
		if booking.Id == excludeBookingId {
			continue
		}

		if booking.SeatNumber == seatNumber {
			return errors.New(fmt.Sprintf("seat number %d is already taken for this round", seatNumber))
		}
	}

	return nil
}

func (s *BookingService) CreateBooking(ctx context.Context, booking *domain.Bookings) (*domain.Bookings, error) {
	// Check if the seat is available
	if err := s.checkSeatAvailability(ctx, booking.RoundId, booking.SeatNumber, ""); err != nil {
		return nil, err
	}

	return s.bookingsRepository.CreateBooking(ctx, booking)
}

func (s *BookingService) GetBookingById(ctx context.Context, id string) (*domain.Bookings, error) {
	return s.bookingsRepository.GetBookingById(ctx, id)
}

func (s *BookingService) GetBookingsByUserId(ctx context.Context, userId string) ([]domain.Bookings, error) {
	return s.bookingsRepository.GetBookingsByUserId(ctx, userId)
}

func (s *BookingService) GetBookingsByRoundId(ctx context.Context, roundId string) ([]domain.Bookings, error) {
	return s.bookingsRepository.GetBookingsByRoundId(ctx, roundId)
}

func (s *BookingService) UpdateBooking(ctx context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error) {
	// Check if the seat is available (excluding the current booking)
	if err := s.checkSeatAvailability(ctx, booking.RoundId, booking.SeatNumber, id); err != nil {
		return nil, err
	}

	return s.bookingsRepository.UpdateBooking(ctx, id, booking)
}

func (s *BookingService) DeleteBooking(ctx context.Context, id string) error {
	return s.bookingsRepository.DeleteBooking(ctx, id)
}
