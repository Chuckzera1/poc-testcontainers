package pet_test

import (
	"context"
	"fmt"
	"os"
	"poc-testcontainers/internal/application"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/pet"
	"poc-testcontainers/internal/repositories/testutils"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB
var repo application.PetRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	testDB, err := testutils.NewTestDatabase(ctx, &models.Pet{})
	if err != nil {
		fmt.Printf("Error getting test db \nReason= %s", err.Error())
		os.Exit(1)
	}
	gormDB := testDB.GormDB
	repo = pet.NewPetRepository(gormDB)
	db = gormDB

	m.Run()

	if err := testDB.Cleanup(ctx); err != nil {
		fmt.Printf("failed to clean up test database: %v\n", err)
	}
}

func cleanUpPetDB(t *testing.T) {
	t.Helper()

	db.Raw("DELETE FROM users")
	db.Raw("DELETE FROM pet")
}
