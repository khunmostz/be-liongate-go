package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/gorm"
)

type GormPerformanceStageRepository struct {
	base *BaseGormRepository
}

func NewGormPerformanceStageRepository(db *gorm.DB) *GormPerformanceStageRepository {
	return &GormPerformanceStageRepository{
		base: NewBaseGormRepository(db),
	}
}

func (r *GormPerformanceStageRepository) GetStages(ctx context.Context) ([]domain.PerformanceStage, error) {
	var stages []domain.PerformanceStage
	if err := r.base.db.WithContext(ctx).Find(&stages).Error; err != nil {
		return nil, err
	}
	return stages, nil
}

func (r *GormPerformanceStageRepository) CreateStage(ctx context.Context, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	// Generate UUID for new stage
	stage.Id = uuid.New().String()

	if err := r.base.Create(ctx, stage); err != nil {
		return nil, err
	}
	return stage, nil
}

func (r *GormPerformanceStageRepository) GetStageById(ctx context.Context, id string) (*domain.PerformanceStage, error) {
	var stage domain.PerformanceStage
	if err := r.base.db.WithContext(ctx).Where("stage_id = ?", id).First(&stage).Error; err != nil {
		return nil, err
	}
	return &stage, nil
}

func (r *GormPerformanceStageRepository) UpdateStage(ctx context.Context, id string, stage *domain.PerformanceStage) (*domain.PerformanceStage, error) {
	existingStage, err := r.GetStageById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.base.db.WithContext(ctx).Model(existingStage).Updates(stage).Error; err != nil {
		return nil, err
	}

	return r.GetStageById(ctx, id)
}

func (r *GormPerformanceStageRepository) DeleteStage(ctx context.Context, id string) error {
	return r.base.db.WithContext(ctx).Where("stage_id = ?", id).Delete(&domain.PerformanceStage{}).Error
}
