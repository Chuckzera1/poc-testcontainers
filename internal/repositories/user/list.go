package user

import (
	"log"
	"poc-testcontainers/internal/models"
)

func (u *userRepository) List(filter *models.User, page int) ([]models.User, error) {
	result := []models.User{}
	err := u.db.
		Where(filter).
		Order("name ASC").
		Limit(10).
		Offset(10 * page).
		Find(&result).
		Error

	if err != nil {
		log.Printf("Error listing users. \nReason= %s", err.Error())
		return nil, err
	}

	return result, err
}
