package dto

type CreatePetReqDTO struct {
	Name              string `json:"name" binding:"required"`
	Age               int    `json:"age" binding:"required"`
	UserResponsibleID uint64 `json:"userResponsibleId" binding:"required"`
}

type CreatePetResDTO struct {
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	UserResponsibleID uint64 `json:"userResponsibleId"`
}

type PetUserResponsibleDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PetListDTO struct {
	ID              uint64                `json:"id"`
	Name            string                `json:"name"`
	Age             int                   `json:"age"`
	UserResponsible PetUserResponsibleDTO `json:"userResponsible"`
}
