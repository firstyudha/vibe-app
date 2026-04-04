# Vibe App

A Go application built with Clean Architecture principles using Gin, GORM, PostgreSQL, and JWT.

## Project Structure

The project follows Clean Architecture principles with the following structure:

```
vibe-app/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   ├── domain/
│   ├── repository/
│   ├── service/
│   ├── delivery/
│   └── utils/
├── pkg/
└── deployments/
    ├── docker/
    └── kubernetes/
```

## Features

- RESTful API with Gin framework
- Database operations with GORM and PostgreSQL
- Authentication with JWT
- Clean Architecture implementation
- Docker support for easy deployment

## Prerequisites

- Go 1.21+
- PostgreSQL
- Docker (optional, for containerization)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/firstyudha/vibe-app.git
   ```

2. Navigate to the project directory:
   ```bash
   cd vibe-app
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

## Usage

### Running Locally

1. Start the PostgreSQL database:
   ```bash
   docker-compose up db
   ```

2. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

### Running with Docker

1. Build and run with docker-compose:
   ```bash
   docker-compose up --build
   ```

## API Endpoints

- `POST /api/v1/register` - Register a new user
- `POST /api/v1/login` - Login and obtain JWT token
- `GET /health` - Health check endpoint

## Environment Variables

- `PORT` - Server port (default: 8080)
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT signing

## Contributing

1. Create a feature branch: `git checkout -b feature/your-feature`
2. Commit your changes: `git commit -am 'Add some feature'`
3. Push to the branch: `git push origin feature/your-feature`
4. Create a new Pull Request

## License

This project is licensed under the MIT License.