package repository

import (
	"context"
)

// BaseRepository defines common CRUD operations
type BaseRepository interface {
	Create(ctx context.Context, entity interface{}) error
	FindByID(ctx context.Context, id string, result interface{}) error
	FindAll(ctx context.Context, filter interface{}, results interface{}) error
	Update(ctx context.Context, id string, entity interface{}) error
	Delete(ctx context.Context, id string) error
}

// ContextWithTimeout adds a timeout to the provided context
func ContextWithTimeout(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, DefaultTimeout)
}

// DefaultTimeout is the default timeout for database operations
const DefaultTimeout = 10 * 1000000000 // 10 seconds
