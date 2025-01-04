package pet

import (
	"net/http"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (l *listPetController) Handle(c *gin.Context) {
	name := c.Query("name")
	page := c.Query("page")

	if page == "" {
		page = "0"
	}

	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "page query params is wrong"})
		return
	}

	result, err := l.repository.List(&model.Pet{
		Name: name,
	}, intPage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to list pets"})
		return
	}

	var dtos []dto.PetListDTO = []dto.PetListDTO{}
	for _, pet := range result {
		dtos = append(dtos, dto.PetListDTO{
			ID:   pet.ID,
			Name: pet.Name,
			Age:  pet.Age,
			UserResponsible: dto.PetUserResponsible{
				ID:   pet.UserResponsible.ID,
				Name: pet.UserResponsible.Name,
				Age:  pet.UserResponsible.Age,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dtos,
	})
}
