package user

import (
	"poc-testcontainers/internal/application"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) application.UserRepository {
	return &userRepository{db}
}
