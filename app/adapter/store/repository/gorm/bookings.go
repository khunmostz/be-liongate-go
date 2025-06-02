package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/gorm"
)

type GormBookingRepository struct {
	db *gorm.DB
}

func NewGormBookingRepository(db *gorm.DB) *GormBookingRepository {
	return &GormBookingRepository{db: db}
}

func (r *GormBookingRepository) CreateBooking(context context.Context, booking *domain.Bookings) (*domain.Bookings, error) {
	// Generate UUID for new booking
	booking.Id = uuid.New().String()

	if err := r.db.WithContext(context).Create(booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *GormBookingRepository) GetBookingById(context context.Context, id string) (*domain.Bookings, error) {
	var booking domain.Bookings
	if err := r.db.WithContext(context).Where("booking_id = ?", id).First(&booking).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *GormBookingRepository) GetBookingsByUserId(context context.Context, userId string) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	if err := r.db.WithContext(context).Where("user_id = ?", userId).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *GormBookingRepository) GetBookingsByRoundId(context context.Context, roundId string) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	if err := r.db.WithContext(context).Where("round_id = ?", roundId).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *GormBookingRepository) UpdateBooking(ctx context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error) {
	existingBooking, err := r.GetBookingById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Model(existingBooking).Updates(booking).Error; err != nil {
		return nil, err
	}

	return r.GetBookingById(ctx, id)
}

func (r *GormBookingRepository) DeleteBooking(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("booking_id = ?", id).Delete(&domain.Bookings{}).Error
}
