package user

import (
	"log"
	"poc-testcontainers/internal/models"
)

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		log.Printf("Error creating user. \nReason= %s", err.Error())
		return nil, err
	}

	return user, nil
}
