package mongo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/common"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	base *BaseMongoRepository
	db   *mongo.Database
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		base: NewBaseMongoRepository(collection),
		db:   collection.Database(),
	}
}

func (r *MongoUserRepository) CreateUser(ctx context.Context, user *domain.Users) (*domain.Users, error) {
	// Generate UUID for new user
	user.Id = uuid.New().String()

	if err := r.base.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoUserRepository) GetUserById(ctx context.Context, id string) (*domain.Users, error) {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"_id": id}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "bookings",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "bookings",
		}}},
	}

	cursor, err := r.base.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.Users
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return &users[0], nil
}

func (r *MongoUserRepository) GetUsersByRole(ctx context.Context, role string) ([]domain.Users, error) {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"role": role}}},
		{{Key: "$lookup", Value: bson.M{
			"from":         "bookings",
			"localField":   "_id",
			"foreignField": "user_id",
			"as":           "bookings",
		}}},
	}

	cursor, err := r.base.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []domain.Users
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MongoUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.Users, error) {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	var user domain.Users
	if err := r.base.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) UpdateUser(ctx context.Context, id string, user *domain.Users) (*domain.Users, error) {
	// First check if user exists
	_, err := r.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Prepare update data
	updateData := bson.M{
		"username": user.Username,
		"password": user.Password,
		"role":     user.Role,
	}

	// Update the user
	if err := r.base.Update(ctx, id, updateData); err != nil {
		return nil, err
	}

	// Return the updated user
	return r.GetUserById(ctx, id)
}

func (r *MongoUserRepository) DeleteUser(ctx context.Context, id string) error {
	return r.base.Delete(ctx, id)
}
