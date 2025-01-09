package pet

import (
	"fmt"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
)

func (c *createPetUsecase) Create(pet *dto.CreatePetReqDTO) (*dto.CreatePetResDTO, error) {
	if pet == nil {
		return nil, fmt.Errorf("invalid input: pet cannot be nil")
	}
	petToCreate := &model.Pet{
		Name:              pet.Name,
		Age:               pet.Age,
		UserResponsibleID: pet.UserResponsibleID,
	}

	petCreated, err := c.repo.Create(petToCreate)
	if err != nil {
		return nil, err
	}

	return &dto.CreatePetResDTO{
		ID:                petCreated.ID,
		Name:              petCreated.Name,
		Age:               petCreated.Age,
		UserResponsibleID: pet.UserResponsibleID,
	}, nil
}
