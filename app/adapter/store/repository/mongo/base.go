package mongo

import (
	"context"
	"errors"
	"reflect"

	"github.com/khunmostz/be-liongate-go/app/adapter/store/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BaseMongoRepository provides a base implementation for MongoDB repositories
type BaseMongoRepository struct {
	collection *mongo.Collection
}

// NewBaseMongoRepository creates a new base MongoDB repository
func NewBaseMongoRepository(collection *mongo.Collection) *BaseMongoRepository {
	return &BaseMongoRepository{
		collection: collection,
	}
}

// Create inserts a new entity into the collection
func (r *BaseMongoRepository) Create(ctx context.Context, entity any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, entity)
	return err
}

// FindByID finds an entity by its ID
func (r *BaseMongoRepository) FindByID(ctx context.Context, id string, result any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("entity not found")
		}
		return err
	}
	return nil
}

// FindAll finds all entities matching the filter
func (r *BaseMongoRepository) FindAll(ctx context.Context, filter any, results any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	// Ensure results is a pointer to a slice
	resultsVal := reflect.ValueOf(results)
	if resultsVal.Kind() != reflect.Ptr || resultsVal.Elem().Kind() != reflect.Slice {
		return errors.New("results parameter must be a pointer to a slice")
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// Update updates an entity by its ID
func (r *BaseMongoRepository) Update(ctx context.Context, id string, entity any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": entity})
	return err
}

// Delete removes an entity by its ID
func (r *BaseMongoRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
