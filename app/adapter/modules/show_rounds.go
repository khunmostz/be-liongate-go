package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"go.uber.org/fx"
)

func ProvideShowRoundsRepository(factory *repository.RepositoryFactory) (port.ShowRoundsRepository, error) {
	return factory.CreateShowRoundRepository()
}

var ShowRoundModule = fx.Options(
	fx.Provide(
		ProvideShowRoundsRepository,
		fx.Annotate(
			services.NewShowRoundService,
			fx.As(new(port.ShowRoundsService)),
		),
		controllers.NewShowRoundsController,
	),
)
