package http

import (
	"vibe-app/internal/delivery/http/handler"
	"vibe-app/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter sets up the Gin router with routes and middleware
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Create handlers
	userHandler := handler.NewUserHandler()

	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes would go here
		// protected.GET("/users", userHandler.GetAll)
		// protected.GET("/users/:id", userHandler.GetByID)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}