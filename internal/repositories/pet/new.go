package pet

import (
	"poc-testcontainers/internal/application"

	"gorm.io/gorm"
)

type petRepository struct {
	db *gorm.DB
}

// Delete implements application.PetRepository.
func (u *petRepository) Delete(ID uint64) error {
	panic("unimplemented")
}

func NewPetRepository(db *gorm.DB) application.PetRepository {
	return &petRepository{db}
}
