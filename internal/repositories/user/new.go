package user

import (
	"poc-testcontainers/internal/application"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Delete implements application.UserRepository.
func (u *userRepository) Delete(ID uint64) error {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) application.UserRepository {
	return &userRepository{db}
}
