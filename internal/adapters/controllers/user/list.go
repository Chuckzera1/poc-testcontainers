package user

import (
	"net/http"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (l *listUserController) Handle(c *gin.Context) {
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

	result, err := l.repository.List(&model.User{
		Name: name,
	}, intPage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to list users"})
		return
	}

	var dtos []dto.UserListDTO = []dto.UserListDTO{}
	for _, user := range result {
		dtos = append(dtos, dto.UserListDTO{
			ID:   user.ID,
			Name: user.Name,
			Age:  user.Age,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dtos,
	})
}
