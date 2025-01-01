package pet_test

import (
	"context"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/pet"
	"poc-testcontainers/internal/repositories/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	ctx := context.Background()
	db, err := testutils.NewTestDatabase(ctx, &models.Pet{})
	if err != nil {
		t.Fatalf("Error getting test db \nReason= %s", err.Error())
	}

	gormDB := db.GormDB
	repo := pet.NewPetRepository(gormDB)
	t.Run("Should create pet correctly", func(t *testing.T) {
		user := models.User{
			Name: "test-name",
			Age:  20,
		}
		gormDB.Create(&user)

		p := models.Pet{
			Name:             "test-pet-name",
			Age:              1,
			UserRespnsibleID: user.ID,
		}
		result, err := repo.Create(&p)

		var petCreated models.Pet
		gormDB.Where("name", "test-pet-name").First(&petCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, petCreated.ID)
		assert.Equal(t, "test-pet-name", petCreated.Name)
		assert.Equal(t, 1, petCreated.Age)
	})
}
