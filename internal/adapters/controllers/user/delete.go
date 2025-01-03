package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *deleteUserController) Handle(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "route param userID is missing"})
		c.Abort()
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "route param userID is not a number"})
		c.Abort()
		return
	}
	err = d.repository.Delete(uint64(userIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusNoContent, "")
}
