package modules

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

// ProvideSwaggerHandler provides the Swagger handler for dependency injection
func ProvideSwaggerHandler() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}

var SwaggerModule = fx.Options(
	fx.Provide(
		ProvideSwaggerHandler,
	),
)
