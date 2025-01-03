package main

import (
	"fmt"
	"log"
	router "poc-testcontainers/internal/adapters/routes"
	"poc-testcontainers/internal/config"
	"poc-testcontainers/internal/database"
	"poc-testcontainers/internal/model"
)

func main() {
	var env config.Env

	err := env.LoadEnv(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	db := database.GetDB(
		env.DATABASE_USER,
		env.DATABASE_PASSWORD,
		env.DATABASE_HOST,
		env.DATABASE_PORT,
		env.DATABASE_NAME,
	)
	db.AutoMigrate(&model.User{}, &model.Pet{})

	r := router.Router(db)

	port := fmt.Sprintf(":%v", env.PORT)
	r.Run(port)
}
