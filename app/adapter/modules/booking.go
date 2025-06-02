package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"go.uber.org/fx"
)

// ProvideBookingsRepository extracts port.BookingsRepository from RepositoryFactory for Fx DI
func ProvideBookingsRepository(factory *repository.RepositoryFactory) (port.BookingsRepository, error) {
	return factory.CreateBookingRepository()
}

var BookingModule = fx.Options(
	fx.Provide(
		ProvideBookingsRepository,
		fx.Annotate(
			services.NewBookingsService,
			fx.As(new(port.BookingsService)),
		),
		controllers.NewBookingsController,
	),
)
