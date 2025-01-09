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
