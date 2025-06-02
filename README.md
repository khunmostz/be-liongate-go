# Liongate Backend API

A backend service for Liongate written in Go using the Gin framework with a clean architecture approach.

## Project Overview

This project is a backend API for Liongate, providing endpoints for managing users, bookings, show rounds, and animals. It follows a hexagonal architecture pattern and uses dependency injection with Uber's FX library.

## Features

- RESTful API with Gin framework
- Swagger documentation
- Support for both MongoDB and PostgreSQL databases
- Clean architecture (Hexagonal/Ports and Adapters)
- Dependency injection with Uber FX
- Containerization with Docker

## Architecture

The project follows a clean architecture approach with the following structure:

```
app/
├── adapter/         # External adapters (controllers, repositories)
├── core/            # Business logic and domain models
│   ├── domain/      # Domain entities
│   ├── port/        # Interface definitions
│   └── services/    # Business logic services
├── cmd/             # Application entry points
└── docs/            # Swagger documentation
```

## Prerequisites

- Go 1.23+
- Docker and Docker Compose
- MongoDB or PostgreSQL

## Getting Started

### Environment Setup

Create a `.env` file in the root directory with the following variables:

```
# Application
APP_ENV=development
SERVER_PORT=8080
DB_TYPE=postgres  # or mongodb

# MongoDB
MONGODB_URI=mongodb://mongo_user:mongo_password@mongodb:27017/liongate?authSource=admin
MONGO_DB=liongate
MONGO_USER=mongo_user
MONGO_PASSWORD=mongo_password
MONGO_PORT=27017

# PostgreSQL
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=liongate
POSTGRES_TIMEZONE=Asia/Bangkok

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here-make-it-long-and-random
JWT_ACCESS_DURATION=15m
JWT_REFRESH_DURATION=168h # 7d
```

### Running with Docker

```bash
docker-compose --env-file .env up -d
```

The API will be available at http://localhost:8080

### Running Locally

```bash
# Install dependencies
go mod download

# Run the application
go run app/cmd/main.go
```

## API Documentation

Swagger documentation is available at `/swagger/index.html` when the application is running.

### Generating Swagger Documentation

To update the Swagger documentation when you make changes to the API:

1. Install the Swag tool:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate the Swagger documentation:
```bash
swag init -g app/cmd/main.go -o app/docs
or
~/go/bin/swag init -g app/cmd/main.go -o app/docs
```


3. Restart the application to see the updated documentation.

## API Endpoints

The API provides the following main endpoints:

- Authentication
- Users management
- Bookings management
- Show rounds management
- Animals management
- Performance stages management

For detailed API documentation, please refer to the Swagger documentation.

## Development

### Hot Reload

This project supports hot reloading using Air. To use it:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Testing

### Running Tests

To run all tests:

```bash
go test ./...
```

To run tests for a specific package:

```bash
go test ./app/core/services/...
```

## License

[MIT](LICENSE)
