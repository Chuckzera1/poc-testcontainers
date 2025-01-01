package main

import (
	"fmt"
	"log"
	"poc-testcontainers/internal/config"
	"poc-testcontainers/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	var env config.Env
	r := gin.Default()

	err := env.LoadEnv(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	database.GetDB(
		env.DATABASE_USER,
		env.DATABASE_HOST,
		env.DATABASE_PORT,
		env.DATABASE_NAME,
	)
	port := fmt.Sprintf(":%v", env.PORT)
	r.Run(port)
}
