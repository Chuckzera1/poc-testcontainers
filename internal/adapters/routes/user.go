package router

import (
	"poc-testcontainers/internal/di/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupUserRouter(db *gorm.DB, r *gin.Engine) *gin.Engine {
	createUser := controllers.DICreateUserController(db)
	listUser := controllers.DIListUserController(db)
	deleteUser := controllers.DIDeleteUserController(db)

	r.POST("/api/v1/user", createUser.Handle)
	r.GET("/api/v1/user", listUser.Handle)
	r.DELETE("/api/v1/user/:id", deleteUser.Handle)

	return r
}
