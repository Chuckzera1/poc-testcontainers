package user

import (
	"poc-testcontainers/internal/application"
)

type createUserController struct {
	repository application.CreateUserRepository
}

func NewCreateUserController(repository application.CreateUserRepository) application.BaseController {
	return &createUserController{
		repository,
	}
}

type listUserController struct {
	repository application.ListUsersRepository
}

func NewListUserController(repository application.ListUsersRepository) application.BaseController {
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
