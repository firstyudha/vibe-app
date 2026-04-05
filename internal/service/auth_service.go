package service

import (
	"vibe-app/internal/domain/user"
	"vibe-app/internal/utils"
)

// AuthService handles authentication logic
type AuthService struct {
	repo       user.UserRepository
	userService *UserService
}

// NewAuthService creates a new AuthService
func NewAuthService(repo user.UserRepository, userService *UserService) *AuthService {
	return &AuthService{
		repo:        repo,
		userService: userService,
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(username, password string) (string, error) {
	// Get user by username
	userEntity, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", &InvalidCredentialsError{}
	}

	// Check password
	if !utils.CheckPassword(password, userEntity.Password) {
		return "", &InvalidCredentialsError{}
	}

	// Generate JWT token
	token, err := utils.GenerateToken(userEntity.ID, userEntity.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// InvalidCredentialsError represents an error for invalid login credentials
type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "invalid credentials"
}