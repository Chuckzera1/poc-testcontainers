package pet

import (
	"log"
	"poc-testcontainers/internal/models"
)

func (u *petRepository) List(filter *models.Pet, page int) ([]models.Pet, error) {
	result := []models.Pet{}
	err := u.db.
		Where(filter).
		Order("name ASC").
		Limit(10).
		Offset(10 * page).
		Find(&result).
		Error

	if err != nil {
		log.Printf("Error listing pets. \nReason= %s", err.Error())
		return nil, err
	}

	return result, err
}
