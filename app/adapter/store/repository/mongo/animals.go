package mongo

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAnimalRepository struct {
	base *BaseMongoRepository
}

func NewMongoAnimalRepository(collection *mongo.Collection) *MongoAnimalRepository {
	return &MongoAnimalRepository{
		base: NewBaseMongoRepository(collection),
	}
}

func (r *MongoAnimalRepository) GetAnimals(ctx context.Context) ([]domain.Animals, error) {
	var animals []domain.Animals
	if err := r.base.FindAll(ctx, bson.M{}, &animals); err != nil {
		return nil, err
	}
	return animals, nil
}

func (r *MongoAnimalRepository) CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error) {
	// Generate UUID for new animal
	animal.Id = uuid.New().String()

	if err := r.base.Create(ctx, animal); err != nil {
		return nil, err
	}
	return animal, nil
}

func (r *MongoAnimalRepository) GetAnimalById(ctx context.Context, id string) (*domain.Animals, error) {
	var animal domain.Animals
	if err := r.base.FindByID(ctx, id, &animal); err != nil {
		return nil, err
	}
	return &animal, nil
}

func (r *MongoAnimalRepository) UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error) {
	// First check if animal exists
	_, err := r.GetAnimalById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare the update data
	updateData := bson.M{
		"name":          animal.Name,
		"species":       animal.Species,
		"type":          animal.Type,
		"show_duration": animal.ShowDuration,
	}

	// Update the animal
	if err := r.base.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Return the updated animal
	return r.GetAnimalById(ctx, id)
}

func (r *MongoAnimalRepository) DeleteAnimal(ctx context.Context, id string) error {
	return r.base.Delete(ctx, id)
}
