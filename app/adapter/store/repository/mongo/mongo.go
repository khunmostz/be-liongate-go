package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/khunmostz/be-liongate-go/app/adapter/config"
	"github.com/khunmostz/be-liongate-go/app/adapter/store/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// InitMongoDB initializes and returns a MongoDB database instance
func InitMongoDB(cfg *config.Config) *mongo.Database {
	ctx, cancel := common.ContextWithTimeout(context.Background())
	defer cancel()

	// Use the MongoDB URI from config if available, otherwise build it
	mongoURI := cfg.MongoDB.URI
	if mongoURI == "" {
		// Construct the MongoDB connection URI if not provided
		host := cfg.MongoDB.Host

		// If we're in a Docker environment and host is localhost, use the service name instead
		if host == "localhost" && os.Getenv("DOCKER_ENV") == "true" {
			host = "mongodb"
		}

		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			cfg.MongoDB.Username,
			cfg.MongoDB.Password,
			host,
			cfg.MongoDB.Port,
			cfg.MongoDB.DbName,
		)
	} else {
		// If URI is provided but contains localhost and we're in Docker, replace with service name
		if os.Getenv("DOCKER_ENV") == "true" && cfg.MongoDB.Host == "localhost" {
			mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
				cfg.MongoDB.Username,
				cfg.MongoDB.Password,
				"mongodb", // Use service name instead of localhost
				cfg.MongoDB.Port,
				cfg.MongoDB.DbName,
			)
		}
	}

	log.Printf("Connecting to MongoDB with URI: %s", mongoURI)

	// Set client options with improved settings
	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetMaxConnIdleTime(30 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to verify connection
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client.Database(cfg.MongoDB.DbName)
}
