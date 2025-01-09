package pet

import (
	"errors"
	"io"
	"net/http"
	"poc-testcontainers/internal/application/dto"

	"github.com/gin-gonic/gin"
)

func (ctrl *createPetController) Handle(c *gin.Context) {
	var reqBody dto.CreatePetReqDTO
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		httpErroMessage := err.Error()
		if errors.Is(err, io.EOF) {
			httpErroMessage = "Request body is required"
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": httpErroMessage})
		c.Abort()
		return
	}

	createdPet, err := ctrl.usecase.Create(&reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create pet"})
		return
	}

	result := &dto.CreatePetResDTO{
		ID:                createdPet.ID,
		Name:              createdPet.Name,
		Age:               createdPet.Age,
		UserResponsibleID: createdPet.UserResponsibleID,
	}

	c.JSON(http.StatusOK, result)
}
