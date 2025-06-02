package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/gorm"
)

type GormAnimalRepository struct {
	base *BaseGormRepository
}

func NewGormAnimalRepository(db *gorm.DB) *GormAnimalRepository {
	return &GormAnimalRepository{
		base: NewBaseGormRepository(db),
	}
}

func (r *GormAnimalRepository) GetAnimals(ctx context.Context) ([]domain.Animals, error) {
	var animals []domain.Animals
	if err := r.base.db.WithContext(ctx).Find(&animals).Error; err != nil {
		return nil, err
	}
	return animals, nil
}

func (r *GormAnimalRepository) CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error) {
	// Generate UUID for new animal
	animal.Id = uuid.New().String()

	if err := r.base.Create(ctx, animal); err != nil {
		return nil, err
	}
	return animal, nil
}

func (r *GormAnimalRepository) GetAnimalById(ctx context.Context, id string) (*domain.Animals, error) {
	var animal domain.Animals
	if err := r.base.db.WithContext(ctx).Where("animal_id = ?", id).First(&animal).Error; err != nil {
		return nil, err
	}
	return &animal, nil
}

func (r *GormAnimalRepository) UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error) {
	existingAnimal, err := r.GetAnimalById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.base.db.WithContext(ctx).Model(existingAnimal).Updates(animal).Error; err != nil {
		return nil, err
	}

	return r.GetAnimalById(ctx, id)
}

func (r *GormAnimalRepository) DeleteAnimal(ctx context.Context, id string) error {
	return r.base.db.WithContext(ctx).Where("animal_id = ?", id).Delete(&domain.Animals{}).Error
}
