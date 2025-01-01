package application

import "poc-testcontainers/internal/models"

type CreatePetRepository interface {
	Create(pet *models.Pet) (*models.Pet, error)
}

type ListPetsRepository interface {
	List(filter *models.Pet, page int) ([]models.Pet, error)
}

type DeletePetRepository interface {
	Delete(ID uint64) error
}

type PetRepository interface {
	CreatePetRepository
	ListPetsRepository
	DeletePetRepository
}
