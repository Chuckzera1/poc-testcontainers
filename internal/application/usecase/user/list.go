package user

import (
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
)

func (l *listUserUseCase) List(name string, page int) ([]dto.UserListResDTO, error) {
	filter := &model.User{
		Name: name,
	}
	users, err := l.repository.List(filter, page)
	if err != nil {
		return nil, err
	}

	results := []dto.UserListResDTO{}

	for _, user := range users {
		results = append(results, dto.UserListResDTO{
			ID:   user.ID,
			Name: user.Name,
			Age:  user.Age,
		})
	}

	return results, nil
}
