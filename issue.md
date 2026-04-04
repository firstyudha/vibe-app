# Project Setup Plan

## Objective
Create a new Go project in the current folder with the following dependencies:
- GORM (ORM library)
- Gin (Web framework)
- PostgreSQL (Database)
- JWT (Authentication)

## High-Level Instructions

### 1. Initialize Go Module
- Run `go mod init vibe-app` in the project root
- This creates go.mod and go.sum files

### 2. Install Dependencies
- `go get -u github.com/gin-gonic/gin` for web framework
- `go get -u gorm.io/gorm` for ORM
- `go get -u gorm.io/driver/postgres` for PostgreSQL driver
- `go get -u github.com/golang-jwt/jwt/v4` for JWT handling

### 3. Project Structure (Clean Architecture)
```
vibe-app/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── domain/
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── auth/
│   │       ├── entity.go
│   │       ├── repository.go
│   │       └── service.go
│   ├── repository/
│   │   ├── postgres/
│   │   │   ├── user_repository.go
│   │   │   └── auth_repository.go
│   │   └── migrations/
│   ├── service/
│   │   ├── auth_service.go
│   │   └── user_service.go
│   ├── delivery/
│   │   └── http/
│   │       ├── handler/
│   │       │   ├── auth_handler.go
│   │       │   └── user_handler.go
│   │       ├── middleware/
│   │       │   └── auth_middleware.go
│   │       └── router.go
│   └── utils/
│       ├── password.go
│       └── jwt.go
├── pkg/
│   └── logger/
└── deployments/
    ├── docker/
    │   ├── Dockerfile
    │   └── docker-compose.yml
    └── kubernetes/
```

### 4. Implementation Steps

#### Project Initialization
- Initialize Go module with `go mod init vibe-app`
- Create the directory structure as outlined above

#### Domain Layer (internal/domain/)
- Define entities (User, Token, etc.) with their methods
- Create repository interfaces that define data operations
- Implement domain services that contain business logic

#### Repository Layer (internal/repository/)
- Implement repository interfaces using PostgreSQL and GORM
- Create database migrations for table creation
- Handle all database interactions through this layer

#### Service Layer (internal/service/)
- Implement business logic that orchestrates between repositories and domain entities
- Auth Service: Handle registration, login, token generation/validation
- User Service: Handle user CRUD operations with proper validation

#### Delivery Layer (internal/delivery/)
- HTTP handlers that receive requests and return responses
- Middleware for authentication, logging, etc.
- Router setup with route definitions and middleware chains

#### Configuration (internal/config/)
- Load configuration from environment variables
- Define configuration structs
- Handle different environments (development, staging, production)

#### Utilities (internal/utils/)
- Password hashing and verification
- JWT token generation and validation
- Other helper functions used across the application

#### Main Application Entry (cmd/server/main.go)
- Initialize configuration
- Setup database connections
- Initialize repositories, services, and handlers
- Start the HTTP server

### 4. Implementation Steps

#### Database Setup (config/database.go)
- Create connection to PostgreSQL using GORM
- Configure connection string with environment variables

#### Models (models/user.go)
- Define User struct with GORM annotations
- Include fields: ID, Username, Email, PasswordHash, CreatedAt, UpdatedAt

#### JWT Utilities (utils/jwt_utils.go)
- Create functions for generating and validating JWT tokens
- Use secure secret key from environment variables

#### Authentication Middleware (middleware/auth_middleware.go)
- Implement middleware to validate JWT tokens
- Extract user information from token

#### Controllers
- Auth Controller: Handle login, register, token refresh
- User Controller: Handle CRUD operations for users

#### Routes (routes/routes.go)
- Define API endpoints
- Apply middleware to protected routes

#### Main Application (main.go)
- Initialize database connection
- Setup Gin router
- Register routes
- Start server

### 5. Environment Variables
Create a `.env` file with:
- DATABASE_URL - PostgreSQL connection string
- JWT_SECRET - Secret key for JWT signing

### 6. Entry Level Developer Notes
- Follow existing code patterns
- All database operations should use GORM
- All API responses should follow consistent format
- Error handling should be explicit
- Validate all input data before processing