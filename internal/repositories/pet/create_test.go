package pet_test

import (
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/pet"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	tx := db.Begin()
	repo := pet.NewPetRepository(tx)

	defer tx.Rollback()

	t.Run("Should create pet correctly", func(t *testing.T) {
		user := models.User{
			Name: "test-name",
			Age:  20,
		}
		tx.Create(&user)

		p := models.Pet{
			Name:             "test-pet-name",
			Age:              1,
			UserRespnsibleID: user.ID,
		}
		result, err := repo.Create(&p)

		var petCreated models.Pet
		tx.Where("name", "test-pet-name").First(&petCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, petCreated.ID)
		assert.Equal(t, "test-pet-name", petCreated.Name)
		assert.Equal(t, 1, petCreated.Age)
	})
}
