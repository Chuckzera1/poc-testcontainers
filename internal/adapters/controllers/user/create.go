package user

import (
	"errors"
	"io"
	"net/http"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"

	"github.com/gin-gonic/gin"
)

func (ctrl *createUserController) Handle(c *gin.Context) {
	var reqBody dto.CreateUserReqBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		httpErroMessage := err.Error()
		if errors.Is(err, io.EOF) {
			httpErroMessage = "Request body is required"
		}
		c.JSON(http.StatusBadRequest, gin.H{"message": httpErroMessage})
		c.Abort()
		return
	}

	um := &model.User{
		Name: reqBody.Name,
		Age:  reqBody.Age,
	}

	createdUser, err := ctrl.repository.Create(um)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   createdUser.ID,
		"name": createdUser.Name,
		"age":  createdUser.Age,
	})
}
