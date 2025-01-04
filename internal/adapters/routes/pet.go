package router

import (
	"poc-testcontainers/internal/di/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupPetRouter(db *gorm.DB, r *gin.Engine) *gin.Engine {
	createPet := controllers.DICreatePetController(db)
	listPet := controllers.DIListPetController(db)

	r.POST("/api/v1/pet", createPet.Handle)
	r.GET("/api/v1/pet", listPet.Handle)

	return r
}
