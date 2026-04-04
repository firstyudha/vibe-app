package main

import (
	"log"

	"vibe-app/internal/config"
	"vibe-app/internal/delivery/http"
	"vibe-app/internal/repository/postgres"
	"vibe-app/internal/service"

	"gorm.io/gorm"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Setup database connection
	db := setupDatabase(cfg)

	// Run migrations
	// runMigrations(db)

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	authRepo := postgres.NewAuthRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo, userService)

	// Initialize handlers and router
	router := http.SetupRouter(db)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupDatabase(cfg *config.Config) *gorm.DB {
	// Setup database connection using GORM
	// This is a placeholder implementation
	// In a real implementation, you would use:
	// db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	// if err != nil {
	//     log.Fatal("Failed to connect to database:", err)
	// }
	// return db
	return nil
}

// func runMigrations(db *gorm.DB) {
// 	// Run migrations
// 	// In a real implementation, you would use:
// 	// db.AutoMigrate(&user.User{})
// }