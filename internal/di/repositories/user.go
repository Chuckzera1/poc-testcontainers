package repositories

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/repositories/user"

	"gorm.io/gorm"
)

func DIUserRepository(db *gorm.DB) application.UserRepository {
	return user.NewUserRepository(db)
}
