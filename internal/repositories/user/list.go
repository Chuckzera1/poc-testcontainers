package user

import (
	"log"
	"poc-testcontainers/internal/model"
)

func (u *userRepository) List(filter *model.User, page int) ([]model.User, error) {
	var result []model.User
	err := u.db.
		Select("id, name, age").
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
