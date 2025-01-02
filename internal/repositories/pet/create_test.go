package pet_test

import (
	"poc-testcontainers/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	cleanUpPetDB(t)
	t.Run("Should create pet correctly", func(t *testing.T) {
		user := models.User{
			Name: "test-name",
			Age:  20,
		}
		db.Create(&user)

		p := models.Pet{
			Name:             "test-pet-name",
			Age:              1,
			UserRespnsibleID: user.ID,
		}
		result, err := repo.Create(&p)

		var petCreated models.Pet
		db.Where("name", "test-pet-name").First(&petCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, petCreated.ID)
		assert.Equal(t, "test-pet-name", petCreated.Name)
		assert.Equal(t, 1, petCreated.Age)
	})
}
