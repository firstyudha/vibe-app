package postgres

import (
	"vibe-app/internal/domain/user"

	"gorm.io/gorm"
)

// UserPostgresRepository implements the user.Repository interface using PostgreSQL
type UserPostgresRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserPostgresRepository
func NewUserRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: db,
	}
}

// Create saves a new user to the database
func (r *UserPostgresRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserPostgresRepository) GetByID(id uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername retrieves a user by username
func (r *UserPostgresRepository) GetByUsername(username string) (*user.User, error) {
	var user user.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserPostgresRepository) GetByEmail(email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update modifies an existing user
func (r *UserPostgresRepository) Update(user *user.User) error {
	return r.db.Save(user).Error
}

// Delete removes a user by ID
func (r *UserPostgresRepository) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}