package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoShowRoundRepository struct {
	base *BaseMongoRepository
}

func NewMongoShowRoundRepository(collection *mongo.Collection) *MongoShowRoundRepository {
	return &MongoShowRoundRepository{
		base: NewBaseMongoRepository(collection),
	}
}

func (r *MongoShowRoundRepository) CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	// Generate UUID for new show round
	showRound.Id = uuid.New().String()

	if err := r.base.Create(ctx, showRound); err != nil {
		return nil, err
	}
	return showRound, nil
}

func (r *MongoShowRoundRepository) GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error) {
	var showRound domain.ShowRounds
	if err := r.base.FindByID(ctx, id, &showRound); err != nil {
		return nil, err
	}
	return &showRound, nil
}

func (r *MongoShowRoundRepository) GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error) {
	var showRounds []*domain.ShowRounds

	cursor, err := r.base.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &showRounds); err != nil {
		return nil, err
	}

	return showRounds, nil
}

func (r *MongoShowRoundRepository) UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	// First check if show round exists
	_, err := r.GetShowRoundById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare update data based on the ShowRounds model fields
	updateData := bson.M{
		"animal_id": showRound.AnimalId,
		"stage_id":  showRound.StageId,
		"show_time": showRound.ShowTime,
	}

	// Update the show round
	if err := r.base.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Return the updated show round
	return r.GetShowRoundById(ctx, id)
}

func (r *MongoShowRoundRepository) DeleteShowRound(ctx context.Context, id string) error {
	return r.base.Delete(ctx, id)
}
