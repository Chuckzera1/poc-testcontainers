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
	usecase application.ListPetUsecase
}

func NewListPetController(usecase application.ListPetUsecase) application.BaseController {
	return &listPetController{
		usecase,
	}
}
