package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"go.uber.org/fx"
)

func ProvidePerformanceStageRepository(factory *repository.RepositoryFactory) (port.PerformanceStageRepository, error) {
	return factory.CreatePerformanceStageRepository()
}

var PerformanceStageModule = fx.Options(

	fx.Provide(
		ProvidePerformanceStageRepository,
		fx.Annotate(
			services.NewPerformanceStageService,
			fx.As(new(port.PerformanceStageService)),
		),
		controllers.NewPerformanceStageController,
	),
)
