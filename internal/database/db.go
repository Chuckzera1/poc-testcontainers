package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func GetDB(DATABASE_USER, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME string) *gorm.DB {
	log.Printf("PASSWORD -> %s", DATABASE_PASSWORD)
	log.Printf("HOST -> %s", DATABASE_USER)

	once.Do(func() {
		db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			DATABASE_HOST,
			DATABASE_PORT,
			DATABASE_USER,
			DATABASE_PASSWORD,
			DATABASE_NAME,
		)), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connection DB: %s", err.Error())
		}

		DB = db
	})

	return DB
}
