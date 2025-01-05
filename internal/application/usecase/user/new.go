package user

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/application/dto"
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

func (l *listUserUseCase) List(name string, page int) ([]dto.UserListResDTO, error) {
	panic("unimplemented")
}

func NewListUserUseCase(repository application.ListUserRepository) application.ListUserUseCase {
	return &listUserUseCase{repository}
}

type deleteUserUseCase struct {
	repository application.DeleteUserRepository
}

func (d *deleteUserUseCase) Delete(id uint64) error {
	panic("unimplemented")
}

func NewDeleteUserUseCase(repository application.DeleteUserRepository) application.DeleteUserUseCase {
	return &deleteUserUseCase{repository}
}
