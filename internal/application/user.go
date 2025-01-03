package application

import "poc-testcontainers/internal/model"

type CreateUserRepository interface {
	Create(user *model.User) (*model.User, error)
}

type ListUsersRepository interface {
	List(filter *model.User, page int) ([]model.User, error)
}

type DeleteUserRepository interface {
	Delete(id uint64) error
}

type UserRepository interface {
	CreateUserRepository
	ListUsersRepository
	DeleteUserRepository
}
