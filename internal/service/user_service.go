package service

import (
	"vibe-app/internal/domain/user"
	"vibe-app/internal/utils"
)

// UserService implements the user.Service interface
type UserService struct {
	repo user.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Register creates a new user
func (s *UserService) Register(request *user.CreateUserRequest) (*user.User, error) {
	// Check if user already exists
	existingUser, _ := s.repo.GetByUsername(request.Username)
	if existingUser != nil {
		return nil, &UserAlreadyExistsError{Username: request.Username}
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	// Create the user entity
	userEntity := &user.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	// Save the user
	err = s.repo.Create(userEntity)
	if err != nil {
		return nil, err
	}

	return userEntity, nil
}

// GetByID retrieves a user by ID
func (s *UserService) GetByID(id uint) (*user.User, error) {
	return s.repo.GetByID(id)
}

// GetByUsername retrieves a user by username
func (s *UserService) GetByUsername(username string) (*user.User, error) {
	return s.repo.GetByUsername(username)
}

// Update updates a user
func (s *UserService) Update(user *user.User) error {
	return s.repo.Update(user)
}

// Delete removes a user
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// UserAlreadyExistsError represents an error when a user already exists
type UserAlreadyExistsError struct {
	Username string
}

func (e *UserAlreadyExistsError) Error() string {
	return "user already exists: " + e.Username
}