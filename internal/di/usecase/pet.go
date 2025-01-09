package usecase

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/application/usecase/pet"
)

func DICreatePetUseCase(repository application.CreatePetRepository) application.CreatePetUsecase {
	return pet.NewCreatePetUseCase(repository)
}
