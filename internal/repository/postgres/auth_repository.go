package postgres

import (
	"vibe-app/internal/domain/user"

	"gorm.io/gorm"
)

// AuthPostgresRepository handles authentication-related database operations
type AuthPostgresRepository struct {
	db *gorm.DB
}

// NewAuthRepository creates a new AuthPostgresRepository
func NewAuthRepository(db *gorm.DB) *AuthPostgresRepository {
	return &AuthPostgresRepository{
		db: db,
	}
}

// GetUserByUsername retrieves a user by username for authentication
func (r *AuthPostgresRepository) GetUserByUsername(username string) (*user.User, error) {
	var user user.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by ID for authentication
func (r *AuthPostgresRepository) GetUserByID(id uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}