package application

import "poc-testcontainers/internal/model"

type CreatePetRepository interface {
	Create(pet *model.Pet) (*model.Pet, error)
}

type ListPetsRepository interface {
	List(filter *model.Pet, page int) ([]model.Pet, error)
}

type DeletePetRepository interface {
	Delete(ID uint64) error
}

type PetRepository interface {
	CreatePetRepository
	ListPetsRepository
	DeletePetRepository
}
