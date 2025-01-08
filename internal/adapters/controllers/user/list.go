package user

import (
	"net/http"
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
		c.Abort()
		return
	}

	result, err := l.useCase.List(name, intPage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to list users"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
