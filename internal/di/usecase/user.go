package usecase

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/application/usecase/user"
)

func DICreateUserUseCase(repository application.CreateUserRepository) application.CreateUserUseCase {
	return user.NewCreateUserUseCase(repository)
}
