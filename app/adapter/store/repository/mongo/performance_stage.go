package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPerformanceStageRepository struct {
	base *BaseMongoRepository
}

func NewMongoPerformanceStageRepository(collection *mongo.Collection) *MongoPerformanceStageRepository {
	return &MongoPerformanceStageRepository{
		base: NewBaseMongoRepository(collection),
	}
}

func (r *MongoPerformanceStageRepository) GetStages(ctx context.Context) ([]domain.PerformanceStage, error) {
	var stages []domain.PerformanceStage
	if err := r.base.FindAll(ctx, bson.M{}, &stages); err != nil {
		return nil, err
	}
	return stages, nil
}

func (r *MongoPerformanceStageRepository) CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	// Generate UUID for new stage
	stage.Id = uuid.New().String()

	if err := r.base.Create(ctx, stage); err != nil {
		return nil, err
	}
	return stage, nil
}

func (r *MongoPerformanceStageRepository) GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error) {
	var stage domain.PerformanceStage
	if err := r.base.FindByID(ctx, id, &stage); err != nil {
		return nil, err
	}
	return &stage, nil
}

func (r *MongoPerformanceStageRepository) UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	// First check if the stage exists
	_, err := r.GetStageById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare the update data
	updateData := bson.M{
		"room_number":    stage.RoomNumber,
		"seat_capacity":  stage.SeatCapacity,
		"price_per_seat": stage.PricePerSeat,
	}

	// Update the stage
	if err := r.base.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Return the updated stage
	return r.GetStageById(ctx, id)
}

func (r *MongoPerformanceStageRepository) DeleteStage(ctx context.Context, id string) error {
	return r.base.Delete(ctx, id)
}
