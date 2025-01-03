package user

import (
	"poc-testcontainers/internal/application"
)

type createUserController struct {
	repository application.CreateUserRepository
}

func NewUserController(repository application.CreateUserRepository) application.BaseController {
	return &createUserController{
		repository,
	}
}
