package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type BookingsRepository interface {
	CreateBooking(context context.Context, booking *domain.Bookings) (*domain.Bookings, error)
	GetBookingById(context context.Context, id string) (*domain.Bookings, error)
	GetBookingsByUserId(context context.Context, userId string) ([]domain.Bookings, error)
	GetBookingsByRoundId(context context.Context, roundId string) ([]domain.Bookings, error)
	UpdateBooking(context context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error)
	DeleteBooking(context context.Context, id string) error
}

type BookingsService interface {
	CreateBooking(context context.Context, booking *domain.Bookings) (*domain.Bookings, error)
	GetBookingById(context context.Context, id string) (*domain.Bookings, error)
	GetBookingsByUserId(context context.Context, userId string) ([]domain.Bookings, error)
	GetBookingsByRoundId(context context.Context, roundId string) ([]domain.Bookings, error)
	UpdateBooking(context context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error)
	DeleteBooking(context context.Context, id string) error
}
