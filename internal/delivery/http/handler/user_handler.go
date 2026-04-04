package handler

import (
	"net/http"

	"vibe-app/internal/domain/user"
	"vibe-app/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	// In a real implementation, this would have service dependencies
}

// NewUserHandler creates a new UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// In a real implementation, this would save to the database
	// For now, we'll just return a success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": map[string]interface{}{
			"username": req.Username,
			"email":    req.Email,
			"password": hashedPassword, // In reality, we wouldn't return this
		},
	})
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real implementation, this would check credentials against the database
	// For now, we'll just generate a token
	token, err := utils.GenerateToken(1, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}