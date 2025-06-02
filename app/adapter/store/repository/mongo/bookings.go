package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoBookingRepository struct {
	base *BaseMongoRepository
}

func NewMongoBookingRepository(collection *mongo.Collection) *MongoBookingRepository {
	return &MongoBookingRepository{
		base: NewBaseMongoRepository(collection),
	}
}

func (r *MongoBookingRepository) CreateBooking(ctx context.Context, booking *domain.Bookings) (*domain.Bookings, error) {
	// Generate UUID for new booking
	booking.Id = uuid.New().String()

	if err := r.base.Create(ctx, booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *MongoBookingRepository) GetBookingById(ctx context.Context, id string) (*domain.Bookings, error) {
	var booking domain.Bookings
	if err := r.base.FindByID(ctx, id, &booking); err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *MongoBookingRepository) GetBookingsByUserId(ctx context.Context, userId string) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	if err := r.base.FindAll(ctx, bson.M{"user_id": userId}, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *MongoBookingRepository) GetBookingsByRoundId(ctx context.Context, roundId string) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	if err := r.base.FindAll(ctx, bson.M{"round_id": roundId}, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *MongoBookingRepository) UpdateBooking(ctx context.Context, id string, booking *domain.Bookings) (*domain.Bookings, error) {
	// First check if booking exists
	_, err := r.GetBookingById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare update data based on the booking model fields
	updateData := bson.M{
		"user_id":     booking.UserId,
		"round_id":    booking.RoundId,
		"seat_number": booking.SeatNumber,
		"price":       booking.Price,
		"qr_code":     booking.QrCode,
	}

	// Update the booking
	if err := r.base.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Return the updated booking
	return r.GetBookingById(ctx, id)
}

func (r *MongoBookingRepository) DeleteBooking(ctx context.Context, id string) error {
	return r.base.Delete(ctx, id)
}
