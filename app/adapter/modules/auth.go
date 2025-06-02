package modules

import (
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"github.com/khunmostz/be-liongate-go/app/core/services"
	"github.com/khunmostz/be-liongate-go/app/utils"
	"go.uber.org/fx"
)

var AuthModule = fx.Options(
	fx.Provide(
		utils.NewJWTService,
		fx.Annotate(
			services.NewAuthService,
			fx.As(new(port.AuthService)),
		),
		controllers.NewAuthController,
	),
)
