package pet

import "poc-testcontainers/internal/application"

type createPetController struct {
	repository application.CreatePetRepository
}

func NewCreatePetController(repository application.CreatePetRepository) application.BaseController {
	return &createPetController{
		repository,
	}
}
