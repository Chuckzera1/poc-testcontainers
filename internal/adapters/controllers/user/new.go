package user

import (
	"poc-testcontainers/internal/application"
)

type createUserController struct {
	useCase application.CreateUserUseCase
}

func NewCreateUserController(useCase application.CreateUserUseCase) application.BaseController {
	return &createUserController{
		useCase,
	}
}

type listUserController struct {
	repository application.ListUserRepository
}

func NewListUserController(repository application.ListUserRepository) application.BaseController {
	return &listUserController{
		repository,
	}
}

type deleteUserController struct {
	repository application.DeleteUserRepository
}

func NewDeleteUserController(repository application.DeleteUserRepository) application.BaseController {
	return &deleteUserController{
		repository,
	}
}
