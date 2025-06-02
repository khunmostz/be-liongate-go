package services

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type AnimalService struct {
	animalRepository port.AnimalsRepository
}

func NewAnimalService(animalRepository port.AnimalsRepository) *AnimalService {
	return &AnimalService{
		animalRepository: animalRepository,
	}
}

func (s *AnimalService) GetAnimals(ctx context.Context) ([]domain.Animals, error) {
	return s.animalRepository.GetAnimals(ctx)
}

func (s *AnimalService) CreateAnimal(ctx context.Context, animal *domain.Animals) (*domain.Animals, error) {
	return s.animalRepository.CreateAnimal(ctx, animal)
}

func (s *AnimalService) GetAnimalById(ctx context.Context, id string) (*domain.Animals, error) {
	return s.animalRepository.GetAnimalById(ctx, id)
}

func (s *AnimalService) UpdateAnimal(ctx context.Context, id string, animal *domain.Animals) (*domain.Animals, error) {
	return s.animalRepository.UpdateAnimal(ctx, id, animal)
}

func (s *AnimalService) DeleteAnimal(ctx context.Context, id string) error {
	return s.animalRepository.DeleteAnimal(ctx, id)
}
