package services

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type PerformanceStageService struct {
	stageRepository port.PerformanceStageRepository
}

func NewPerformanceStageService(stageRepository port.PerformanceStageRepository) *PerformanceStageService {
	return &PerformanceStageService{
		stageRepository: stageRepository,
	}
}

func (s *PerformanceStageService) GetStages(ctx context.Context) ([]domain.PerformanceStage, error) {
	return s.stageRepository.GetStages(ctx)
}

func (s *PerformanceStageService) CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	return s.stageRepository.CreateStage(ctx, stage)
}

func (s *PerformanceStageService) GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error) {
	return s.stageRepository.GetStageById(ctx, id)
}

func (s *PerformanceStageService) UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	return s.stageRepository.UpdateStage(ctx, id, stage)
}

func (s *PerformanceStageService) DeleteStage(ctx context.Context, id string) error {
	return s.stageRepository.DeleteStage(ctx, id)
}
