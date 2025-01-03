package user

import (
	"log"
	"poc-testcontainers/internal/model"
)

func (u *userRepository) Delete(id uint64) error {
	err := u.db.Delete(&model.User{ID: id}).Error
	if err != nil {
		log.Printf("Error creating user. \nReason= %s", err.Error())
		return err
	}

	return nil
}
