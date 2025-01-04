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

type listPetController struct {
	repository application.ListPetsRepository
}

func NewListPetController(repository application.ListPetsRepository) application.BaseController {
	return &listPetController{
		repository,
	}
}
