package services

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
	"github.com/khunmostz/be-liongate-go/app/core/port"
)

type ShowRoundService struct {
	showRoundRepository port.ShowRoundsRepository
}

func NewShowRoundService(showRoundRepository port.ShowRoundsRepository) *ShowRoundService {
	return &ShowRoundService{
		showRoundRepository: showRoundRepository,
	}
}

func (s *ShowRoundService) CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	return s.showRoundRepository.CreateShowRound(ctx, showRound)
}

func (s *ShowRoundService) GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error) {
	return s.showRoundRepository.GetShowRoundById(ctx, id)
}

func (s *ShowRoundService) GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error) {
	return s.showRoundRepository.GetAllShowRounds(ctx)
}

func (s *ShowRoundService) UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error) {
	return s.showRoundRepository.UpdateShowRound(ctx, id, showRound)
}

func (s *ShowRoundService) DeleteShowRound(ctx context.Context, id string) error {
	return s.showRoundRepository.DeleteShowRound(ctx, id)
}
