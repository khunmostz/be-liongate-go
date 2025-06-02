package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khunmostz/be-liongate-go/app/adapter/config"
	"github.com/khunmostz/be-liongate-go/app/adapter/controllers"
	"github.com/khunmostz/be-liongate-go/app/adapter/modules"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/repository"
	GormStore "github.com/khunmostz/be-liongate-go/app/adapter/store/repository/gorm"
	MongoStore "github.com/khunmostz/be-liongate-go/app/adapter/store/repository/mongo"
	_ "github.com/khunmostz/be-liongate-go/app/docs" // Import swagger docs
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// @title           Liongate API
// @version         1.0
// @description     A Liongate service API in Go using Gin framework.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func NewRouter() *gin.Engine {
	return gin.Default()
}

func NewConfig() *config.Config {
	return config.NewConfig()
}

func NewRepositoryImpl(cfg *config.Config) (*repository.RepositoryFactory, error) {
	var mongoDb *mongo.Database
	var postgresDb *gorm.DB

	if cfg.IsMongoDB() {
		mongoDb = MongoStore.InitMongoDB(cfg)
	} else if cfg.IsPostgres() {
		postgresDb = GormStore.InitPostgresDB(cfg)
	}

	return repository.NewRepositoryFactory(cfg, mongoDb, postgresDb), nil
}

func RegisterRoutes(
	lc fx.Lifecycle,
	router *gin.Engine,
	authController *controllers.AuthController,
	userController *controllers.UsersController,
	bookingController *controllers.BookingsController,
	showRoundController *controllers.ShowRoundsController,
	animalController *controllers.AnimalsController,
	performanceStageController *controllers.PerformanceStageController,
	swaggerHandler gin.HandlerFunc,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Swagger documentation endpoint
			router.GET("/swagger/*any", swaggerHandler)

			authController.RegisterRoutes(router)
			userController.RegisterRoutes(router)
			bookingController.RegisterRoutes(router)
			showRoundController.RegisterRoutes(router)
			animalController.RegisterRoutes(router)
			performanceStageController.RegisterRoutes(router)

			go router.Run(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			NewRouter,
			NewConfig,
			NewRepositoryImpl,
		),
		modules.AuthModule,
		modules.UserModule,
		modules.BookingModule,
		modules.SwaggerModule,
		modules.ShowRoundModule,
		modules.AnimalModule,
		modules.PerformanceStageModule,
		fx.Invoke(RegisterRoutes),
	)

	app.Run()
}
