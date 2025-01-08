package usecase

import (
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/application/usecase/user"
)

func DICreateUserUseCase(repository application.CreateUserRepository) application.CreateUserUseCase {
	return user.NewCreateUserUseCase(repository)
}

func DIListUserUseCase(repository application.ListUserRepository) application.ListUserUseCase {
	return user.NewListUserUseCase(repository)
}

func DIDeleteUserUseCase(repository application.DeleteUserRepository) application.DeleteUserUseCase {
	return user.NewDeleteUserUseCase(repository)
}
