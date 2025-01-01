package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func GetDB(DATABASE_USER, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME string) *gorm.DB {
	sync.OnceFunc(func() {
		db, err := gorm.Open(postgres.Open(fmt.Sprintf("Conectando ao banco de dados: %s@%s:%s/%s\n",
			DATABASE_USER,
			DATABASE_HOST,
			DATABASE_PORT,
			DATABASE_NAME,
		)), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connection DB: %s", err.Error())
		}

		DB = db
	})

	return DB
}
