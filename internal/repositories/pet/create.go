package pet

import (
	"errors"
	"fmt"
	"log"
	"poc-testcontainers/internal/model"

	"gorm.io/gorm"
)

func (u *petRepository) Create(pet *model.Pet) (*model.Pet, error) {
	var user model.User
	err := u.db.Where("id", pet.UserResponsibleID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user responsible does not exist or has been deleted")
		}
		return nil, err
	}

	err = u.db.Create(pet).Error
	if err != nil {
		log.Printf("Error creating pet. \nReason= %s", err.Error())
		return nil, err
	}

	return pet, err
}
