package gorm

import (
	"context"
	"errors"
	"reflect"

	"github.com/khunmostz/be-liongate-go/app/adapter/store/common"
	"gorm.io/gorm"
)

// BaseGormRepository provides a base implementation for GORM repositories
type BaseGormRepository struct {
	db *gorm.DB
}

// NewBaseGormRepository creates a new base GORM repository
func NewBaseGormRepository(db *gorm.DB) *BaseGormRepository {
	return &BaseGormRepository{
		db: db,
	}
}

// Create inserts a new entity into the database
func (r *BaseGormRepository) Create(ctx context.Context, entity any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	return r.db.WithContext(ctx).Create(entity).Error
}

// FindByID finds an entity by its ID
func (r *BaseGormRepository) FindByID(ctx context.Context, id string, result any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	err := r.db.WithContext(ctx).First(result, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("entity not found")
		}
		return err
	}
	return nil
}

// FindAll finds all entities matching the conditions
func (r *BaseGormRepository) FindAll(ctx context.Context, conditions any, results any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	// Ensure results is a pointer to a slice
	resultsVal := reflect.ValueOf(results)
	if resultsVal.Kind() != reflect.Ptr || resultsVal.Elem().Kind() != reflect.Slice {
		return errors.New("results parameter must be a pointer to a slice")
	}

	return r.db.WithContext(ctx).Where(conditions).Find(results).Error
}

// Update updates an entity by its ID
func (r *BaseGormRepository) Update(ctx context.Context, id string, entity any) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	return r.db.WithContext(ctx).Model(entity).Where("id = ?", id).Updates(entity).Error
}

// Delete removes an entity by its ID
func (r *BaseGormRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := common.ContextWithTimeout(ctx)
	defer cancel()

	// The entity type must be provided for GORM to know which table to delete from
	// This is a limitation of this generic approach with GORM
	// In the specific repositories, you'll need to provide the correct entity type
	return errors.New("delete operation must be implemented in specific repositories")
}
