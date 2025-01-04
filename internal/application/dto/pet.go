package dto

type CreatePetReqBody struct {
	Name              string `json:"name" binding:"required"`
	Age               int    `json:"age" binding:"required"`
	UserResponsibleID uint64 `json:"userResponsibleId" binding:"required"`
}

type CreatePetResBody struct {
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	UserResponsibleID uint64 `json:"userResponsibleId"`
}

type PetUserResponsible struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PetListDTO struct {
	ID              uint64             `json:"id"`
	Name            string             `json:"name"`
	Age             int                `json:"age"`
	UserResponsible PetUserResponsible `json:"userResponsible"`
}
