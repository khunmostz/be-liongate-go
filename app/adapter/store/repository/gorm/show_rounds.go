package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"gorm.io/gorm"
)

type GormShowRoundRepository struct {
	base *BaseGormRepository
}

func NewGormShowRoundRepository(db *gorm.DB) *GormShowRoundRepository {
	return &GormShowRoundRepository{
		base: NewBaseGormRepository(db),
	}
}

func (r *GormShowRoundRepository) CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	// Generate UUID for new show round
	showRound.Id = uuid.New().String()

	if err := r.base.Create(ctx, showRound); err != nil {
		return nil, err
	}
	return showRound, nil
}

func (r *GormShowRoundRepository) GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error) {
	var showRound domain.ShowRounds
	if err := r.base.db.WithContext(ctx).Where("round_id = ?", id).First(&showRound).Error; err != nil {
		return nil, err
	}
	return &showRound, nil
}

func (r *GormShowRoundRepository) GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error) {
	var showRounds []*domain.ShowRounds
	if err := r.base.db.WithContext(ctx).Find(&showRounds).Error; err != nil {
		return nil, err
	}
	return showRounds, nil
}

func (r *GormShowRoundRepository) UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	existingShowRound, err := r.GetShowRoundById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.base.db.WithContext(ctx).Model(existingShowRound).Updates(showRound).Error; err != nil {
		return nil, err
	}

	return r.GetShowRoundById(ctx, id)
}

func (r *GormShowRoundRepository) DeleteShowRound(ctx context.Context, id string) error {
	return r.base.db.WithContext(ctx).Where("round_id = ?", id).Delete(&domain.ShowRounds{}).Error
}
