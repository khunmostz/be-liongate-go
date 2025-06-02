package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"go.uber.org/fx"
)

// ProvideUsersRepository extracts port.UsersRepository from RepositoryFactory for Fx DI
func ProvideUsersRepository(factory *repository.RepositoryFactory) (port.UsersRepository, error) {
	return factory.CreateUserRepository()
}

var UserModule = fx.Options(
	fx.Provide(
		ProvideUsersRepository,
		fx.Annotate(
			services.NewUsersService,
			fx.As(new(port.UsersService)),
		),
		controllers.NewUsersController,
	),
)
