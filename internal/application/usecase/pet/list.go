package pet

import (
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
)

func (l *listPetUsecase) List(name string, page int) ([]*dto.PetListDTO, error) {
	result, err := l.repo.List(&model.Pet{
		Name: name,
	}, page)
	if err != nil {
		return nil, err
	}

	var dtos []*dto.PetListDTO = []*dto.PetListDTO{}
	for _, pet := range result {
		dtos = append(dtos, &dto.PetListDTO{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
			UserResponsible: &dto.PetUserResponsibleDTO{
				ID:   pet.UserResponsible.ID,
				Name: pet.UserResponsible.Name,
				Age:  pet.UserResponsible.Age,
			},
		})
	}

	return dtos, nil
}
