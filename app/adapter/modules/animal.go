package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"go.uber.org/fx"
)

func ProvideAnimalsRepository(factory *repository.RepositoryFactory) (port.AnimalsRepository, error) {
	return factory.CreateAnimalRepository()
}

var AnimalModule = fx.Options(
	fx.Provide(
		ProvideAnimalsRepository,
		fx.Annotate(
			services.NewAnimalService,
			fx.As(new(port.AnimalsService)),
		),
		controllers.NewAnimalsController,
	),
	fx.Invoke(func(controller *controllers.AnimalsController) {}),
)
