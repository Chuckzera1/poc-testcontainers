package pet

import (
	"log"
	"poc-testcontainers/internal/model"
)

func (u *petRepository) List(filter *model.Pet, page int) ([]model.Pet, error) {
	result := []model.Pet{}
	err := u.db.
		Where(filter).
		Preload("UserResponsible").
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
