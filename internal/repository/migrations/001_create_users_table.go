package migrations

import (
	"gorm.io/gorm"
)

// UserMigration represents the users table migration
func UserMigration(db *gorm.DB) error {
	// In a real implementation, this would create the users table
	// For now, we'll just return nil
	return nil
}