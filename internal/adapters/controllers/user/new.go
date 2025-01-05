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
	usecase application.ListUserUseCase
}

func NewListUserController(usecase application.ListUserUseCase) application.BaseController {
	return &listUserController{
		usecase,
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
