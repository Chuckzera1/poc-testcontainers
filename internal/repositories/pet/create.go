package pet

import (
	"log"
	"poc-testcontainers/internal/models"
)

func (u *petRepository) Create(pet *models.Pet) (*models.Pet, error) {
	err := u.db.Create(pet).Error
	if err != nil {
		log.Printf("Error creating pet. \nReason= %s", err.Error())
		return nil, err
	}

	return pet, err
}
