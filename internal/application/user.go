package application

import (
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
)

type CreateUserRepository interface {
	Create(user *model.User) (*model.User, error)
}

type ListUserRepository interface {
	List(filter *model.User, page int) ([]model.User, error)
}

type DeleteUserRepository interface {
	Delete(id uint64) error
}

type UserRepository interface {
	CreateUserRepository
	ListUserRepository
	DeleteUserRepository
}

type CreateUserUseCase interface {
	Create(user *dto.CreateUserReqDTO) (*dto.CreateUserResDTO, error)
}

type ListUserUseCase interface {
	List(name string, page int) ([]dto.UserListResDTO, error)
}

type DeleteUserUseCase interface {
	Delete(id uint64) error
}
