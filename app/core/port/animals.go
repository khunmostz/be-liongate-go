package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type AnimalsRepository interface {
	GetAnimals(ctx context.Context) ([]domain.Animals, error)
	CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error)
	GetAnimalById(ctx context.Context, id string) (*domain.Animals, error)
	UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error)
	DeleteAnimal(ctx context.Context, id string) error
}

type AnimalsService interface {
	GetAnimals(ctx context.Context) ([]domain.Animals, error)
	CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error)
	GetAnimalById(ctx context.Context, id string) (*domain.Animals, error)
	UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error)
	DeleteAnimal(ctx context.Context, id string) error
}
