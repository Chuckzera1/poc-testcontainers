package user

import (
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
)

func (c *createUserUseCase) Create(user *dto.CreateUserReqDTO) (*dto.CreateUserResDTO, error) {
	userToCreate := model.User{
		Name: user.Name,
		Age:  user.Age,
	}

	u, err := c.repository.Create(&userToCreate)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResDTO{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}, nil
}
