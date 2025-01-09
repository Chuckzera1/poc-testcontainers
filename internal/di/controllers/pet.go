package controllers

import (
	"poc-testcontainers/internal/adapters/controllers/pet"
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/di/repositories"
	"poc-testcontainers/internal/di/usecase"

	"gorm.io/gorm"
)

func DICreatePetController(db *gorm.DB) application.BaseController {
	repo := repositories.DIPetRepository(db)
	usecase := usecase.DICreatePetUseCase(repo)

	return pet.NewCreatePetController(usecase)
}

func DIListPetController(db *gorm.DB) application.BaseController {
	repo := repositories.DIPetRepository(db)
	usecase := usecase.DIListPetUseCase(repo)

	return pet.NewListPetController(usecase)
}
