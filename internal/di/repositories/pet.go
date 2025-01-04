package repositories

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/repositories/pet"

	"gorm.io/gorm"
)

func DIPetRepository(db *gorm.DB) application.PetRepository {
	return pet.NewPetRepository(db)
}
