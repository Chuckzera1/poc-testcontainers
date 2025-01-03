package application

import "github.com/gin-gonic/gin"

type BaseController interface {
	Handle(c *gin.Context)
}
