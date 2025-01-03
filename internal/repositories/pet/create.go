package pet

import (
	"log"
	"poc-testcontainers/internal/model"
)

func (u *petRepository) Create(pet *model.Pet) (*model.Pet, error) {
	err := u.db.Create(pet).Error
	if err != nil {
		log.Printf("Error creating pet. \nReason= %s", err.Error())
		return nil, err
	}

	return pet, err
}
