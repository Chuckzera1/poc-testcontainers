package application

import "poc-testcontainers/internal/models"

type CreateUserRepository interface {
	Create(user *models.User) (*models.User, error)
}

type ListUsersRepository interface {
	List(filter *models.User, page int) ([]models.User, error)
}

type DeleteUserRepository interface {
	Delete(id uint64) error
}

type UserRepository interface {
	CreateUserRepository
	ListUsersRepository
	DeleteUserRepository
}
