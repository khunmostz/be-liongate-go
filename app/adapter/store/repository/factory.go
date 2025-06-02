package repository

import (
	"fmt"

	"github.com/khunmostz/be-liongate-go/app/adapter/config"
	localGorm "github.com/khunmostz/be-liongate-go/app/adapter/store/repository/gorm"
	localMongo "github.com/khunmostz/be-liongate-go/app/adapter/store/repository/mongo"
	"github.com/khunmostz/be-liongate-go/app/core/port"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// RepositoryFactory creates repositories based on configuration
type RepositoryFactory struct {
	config     *config.Config
	mongoDB    *mongo.Database
	postgresql *gorm.DB
}

// NewRepositoryFactory creates a new repository factory
func NewRepositoryFactory(cfg *config.Config, mongoDb *mongo.Database, postgresDb *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{
		config:     cfg,
		mongoDB:    mongoDb,
		postgresql: postgresDb,
	}
}

// CreateUserRepository returns the appropriate user repository implementation
func (f *RepositoryFactory) CreateUserRepository() (port.UsersRepository, error) {
	switch f.config.Database.DbType {
	case "mongodb":
		if f.mongoDB == nil {
			return nil, fmt.Errorf("mongodb connection is not initialized")
		}
		collection := f.mongoDB.Collection("users")
		return localMongo.NewMongoUserRepository(collection), nil
	case "postgresql":
		if f.postgresql == nil {
			return nil, fmt.Errorf("postgresql connection is not initialized")
		}
		return localGorm.NewGormUsersRepository(f.postgresql), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.Database.DbType)
	}
}

// CreateBookingRepository returns the appropriate booking repository implementation
func (f *RepositoryFactory) CreateBookingRepository() (port.BookingsRepository, error) {
	switch f.config.Database.DbType {
	case "mongodb":
		if f.mongoDB == nil {
			return nil, fmt.Errorf("mongodb connection is not initialized")
		}
		collection := f.mongoDB.Collection("bookings")
		return localMongo.NewMongoBookingRepository(collection), nil
	case "postgresql":
		if f.postgresql == nil {
			return nil, fmt.Errorf("postgresql connection is not initialized")
		}
		return localGorm.NewGormBookingRepository(f.postgresql), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.Database.DbType)
	}
}

func (f *RepositoryFactory) CreateShowRoundRepository() (port.ShowRoundsRepository, error) {
	switch f.config.Database.DbType {
	case "mongodb":
		if f.mongoDB == nil {
			return nil, fmt.Errorf("mongodb connection is not initialized")
		}
		collection := f.mongoDB.Collection("show_rounds")
		return localMongo.NewMongoShowRoundRepository(collection), nil
	case "postgresql":
		if f.postgresql == nil {
			return nil, fmt.Errorf("postgresql connection is not initialized")
		}
		return localGorm.NewGormShowRoundRepository(f.postgresql), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.Database.DbType)
	}
}

func (f *RepositoryFactory) CreateAnimalRepository() (port.AnimalsRepository, error) {
	switch f.config.Database.DbType {
	case "mongodb":
		if f.mongoDB == nil {
			return nil, fmt.Errorf("mongodb connection is not initialized")
		}
		collection := f.mongoDB.Collection("animals")
		return localMongo.NewMongoAnimalRepository(collection), nil
	case "postgresql":
		if f.postgresql == nil {
			return nil, fmt.Errorf("postgresql connection is not initialized")
		}
		return localGorm.NewGormAnimalRepository(f.postgresql), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.Database.DbType)
	}
}

func (f *RepositoryFactory) CreatePerformanceStageRepository() (port.PerformanceStageRepository, error) {
	switch f.config.Database.DbType {
	case "mongodb":
		if f.mongoDB == nil {
			return nil, fmt.Errorf("mongodb connection is not initialized")
		}
		collection := f.mongoDB.Collection("performance_stages")
		return localMongo.NewMongoPerformanceStageRepository(collection), nil
	case "postgresql":
		if f.postgresql == nil {
			return nil, fmt.Errorf("postgresql connection is not initialized")
		}
		return localGorm.NewGormPerformanceStageRepository(f.postgresql), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.Database.DbType)
	}
}
