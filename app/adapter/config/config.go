package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	DbType   string // "mongodb" or "postgresql"
	SSLMode  string
}

type Config struct {
	Server   Server
	Database Database
	MongoDB  MongoDBConfig
	Postgres PostgresConfig
	Env      string
}

type MongoDBConfig struct {
	URI      string
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
	TimeZone string
}

// LoadEnv loads the environment configuration based on the environment
func LoadEnv(env string) error {
	if env == "" {
		env = "dev" // default to dev environment
	}

	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("error loading %s file: %v", envFile, err)
	}
	fmt.Printf("Loaded environment: %s\n", env)

	return nil
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	env := os.Getenv("APP_ENV")
	if err := LoadEnv(env); err != nil {
		fmt.Printf("Warning: %v\n", err)
	}

	dbType := getEnv("DB_TYPE", "postgresql")
	var dbConfig Database
	switch dbType {
	case "mongodb":
		dbConfig = Database{
			Host:     getEnv("MONGO_HOST", getEnv("DB_HOST", "localhost")),
			Port:     getEnv("MONGO_PORT", getEnv("DB_PORT", "27017")),
			Username: getEnv("MONGO_USER", getEnv("DB_USERNAME", "")),
			Password: getEnv("MONGO_PASSWORD", getEnv("DB_PASSWORD", "")),
			DbName:   getEnv("MONGO_DB", getEnv("DB_NAME", "liongate")),
			DbType:   getEnv("DB_TYPE", "mongodb"),
			SSLMode:  getEnv("MONGO_SSL_MODE", getEnv("DB_SSL_MODE", "disable")),
		}
	default:
		dbConfig = Database{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			Username: getEnv("POSTGRES_USER", ""),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DbName:   getEnv("POSTGRES_DB", "liongate"),
			DbType:   getEnv("DB_TYPE", "postgresql"),
			SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		}
	}

	mongoConfig := MongoDBConfig{
		URI:      getEnv("MONGODB_URI", ""),
		Host:     getEnv("MONGO_HOST", getEnv("DB_HOST", "localhost")),
		Port:     getEnv("MONGO_PORT", getEnv("DB_PORT", "27017")),
		Username: getEnv("MONGO_USER", getEnv("DB_USERNAME", "")),
		Password: getEnv("MONGO_PASSWORD", getEnv("DB_PASSWORD", "")),
		DbName:   getEnv("MONGO_DB", getEnv("DB_NAME", "liongate")),
	}

	postgresConfig := PostgresConfig{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		Port:     getEnv("POSTGRES_PORT", "5432"),
		Username: getEnv("POSTGRES_USER", ""),
		Password: getEnv("POSTGRES_PASSWORD", ""),
		DbName:   getEnv("POSTGRES_DB", "liongate"),
		SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		TimeZone: getEnv("POSTGRES_TIMEZONE", "UTC"),
	}

	return &Config{
		Env: env,
		Server: Server{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: dbConfig,
		MongoDB:  mongoConfig,
		Postgres: postgresConfig,
	}
}

// getEnv retrieves an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// GetDatabaseConfig returns the appropriate database configuration
func (c *Config) GetDatabaseConfig() any {
	switch c.Database.DbType {
	case "mongodb":
		return c.MongoDB
	case "postgresql":
		return c.Postgres
	default:
		return nil
	}
}

// IsMongoDB checks if MongoDB is the selected database
func (c *Config) IsMongoDB() bool {
	return c.Database.DbType == "mongodb"
}

// IsPostgres checks if PostgreSQL is the selected database
func (c *Config) IsPostgres() bool {
	return c.Database.DbType == "postgresql"
}
