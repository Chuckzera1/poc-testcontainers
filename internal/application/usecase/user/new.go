package user

import (
	"poc-testcontainers/internal/application"
)

type createUserUseCase struct {
	repository application.CreateUserRepository
}

func NewCreateUserUseCase(repository application.CreateUserRepository) application.CreateUserUseCase {
	return &createUserUseCase{repository}
}

type listUserUseCase struct {
	repository application.ListUserRepository
}

func NewListUserUseCase(repository application.ListUserRepository) application.ListUserUseCase {
	return &listUserUseCase{repository}
}

type deleteUserUseCase struct {
	repository application.DeleteUserRepository
}

func NewDeleteUserUseCase(repository application.DeleteUserRepository) application.DeleteUserUseCase {
	return &deleteUserUseCase{repository}
}
