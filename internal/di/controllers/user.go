package controllers

import (
	"poc-testcontainers/internal/adapters/controllers/user"
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/di/repositories"

	"gorm.io/gorm"
)

func DICreateUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)

	return user.NewCreateUserController(repo)
}

func DIListUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)

	return user.NewListUserController(repo)
}

func DIDeleteUserController(db *gorm.DB) application.BaseController {
	repo := repositories.DIUserRepository(db)

	return user.NewDeleteUserController(repo)
}
