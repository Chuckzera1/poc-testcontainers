package pet

import (
	"poc-testcontainers/internal/application"
)

type createPetUsecase struct {
	repo application.CreatePetRepository
}

func NewCreatePetUseCase(repo application.CreatePetRepository) application.CreatePetUsecase {
	return &createPetUsecase{repo}
}

type listPetUsecase struct {
	repo application.ListPetsRepository
}

func NewListPetUseCase(repo application.ListPetsRepository) application.ListPetUsecase {
	return &listPetUsecase{repo}
}
