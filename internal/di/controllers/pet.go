package controllers

import (
	"poc-testcontainers/internal/adapters/controllers/pet"
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/di/repositories"

	"gorm.io/gorm"
)

func DICreatePetController(db *gorm.DB) application.BaseController {
	repo := repositories.DIPetRepository(db)

	return pet.NewCreatePetController(repo)
}

func DIListPetController(db *gorm.DB) application.BaseController {
	repo := repositories.DIPetRepository(db)

	return pet.NewListPetController(repo)
}
