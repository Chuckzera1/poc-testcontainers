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
	useCase application.ListUserUseCase
}

func NewListUserController(useCase application.ListUserUseCase) application.BaseController {
	return &listUserController{
		useCase,
	}
}

type deleteUserController struct {
	useCase application.DeleteUserUseCase
}

func NewDeleteUserController(useCase application.DeleteUserUseCase) application.BaseController {
	return &deleteUserController{
		useCase,
	}
}
