package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type PerformanceStageRepository interface {
	GetStages(ctx context.Context) ([]domain.PerformanceStage, error)
	CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error)
	GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error)
	UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error)
	DeleteStage(ctx context.Context, id string) error
}

type PerformanceStageService interface {
	GetStages(ctx context.Context) ([]domain.PerformanceStage, error)
	CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error)
	GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error)
	UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error)
	DeleteStage(ctx context.Context, id string) error
}
