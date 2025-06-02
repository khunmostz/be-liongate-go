package port

import (
	"context"

	"github.com/khunmostz/be-liongate-go/app/core/domain"
)

type ShowRoundsRepository interface {
	CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error)
	GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error)
	GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error)
	UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error)
	DeleteShowRound(ctx context.Context, id string) error
}

type ShowRoundsService interface {
	CreateShowRound(ctx context.Context, showRound *domain.ShowRounds) (*domain.ShowRounds, error)
	GetShowRoundById(ctx context.Context, id string) (*domain.ShowRounds, error)
	GetAllShowRounds(ctx context.Context) ([]*domain.ShowRounds, error)
	UpdateShowRound(ctx context.Context, id string, showRound *domain.ShowRounds) (*domain.ShowRounds, error)
	DeleteShowRound(ctx context.Context, id string) error
}
