package pet

import "poc-testcontainers/internal/application"

type createPetController struct {
	usecase application.CreatePetUsecase
}

func NewCreatePetController(usecase application.CreatePetUsecase) application.BaseController {
	return &createPetController{
		usecase,
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
