package controllers

import (
	"poc-testcontainers/internal/adapters/controllers/user"
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/di/repositories"
	"poc-testcontainers/internal/di/usecase"

	"gorm.io/gorm"
)

func DICreateUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)
	usecase := usecase.DICreateUserUseCase(repo)

	return user.NewCreateUserController(usecase)
}

func DIListUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)
	usecase := usecase.DIListUserUseCase(repo)

	return user.NewListUserController(usecase)
}

func DIDeleteUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)
	u := usecase.DIDeleteUserUseCase(repo)

	return user.NewDeleteUserController(u)
}
