package pet_test

import (
	"poc-testcontainers/internal/model"
	"poc-testcontainers/internal/repositories/pet"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	t.Run("Should create pet correctly", func(t *testing.T) {
		tx := db.Begin()
		repo := pet.NewPetRepository(tx)

		defer tx.Rollback()

		user := model.User{
			Name: "test-name",
			Age:  20,
		}
		tx.Create(&user)

		p := model.Pet{
			Name:              "test-pet-name",
			Age:               1,
			UserResponsibleID: user.ID,
		}
		result, err := repo.Create(&p)

		var petCreated model.Pet
		tx.Where("name", "test-pet-name").First(&petCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, petCreated.ID)
		assert.Equal(t, "test-pet-name", petCreated.Name)
		assert.Equal(t, 1, petCreated.Age)
	})

	t.Run("Should not create pet for user deleted", func(t *testing.T) {
		tx := db.Begin()
		repo := pet.NewPetRepository(tx)

		defer tx.Rollback()

		user := model.User{
			Name: "test-name",
			Age:  20,
		}
		tx.Create(&user)
		tx.Delete(&model.User{ID: user.ID})

		p := model.Pet{
			Name:              "test-pet-name",
			Age:               1,
			UserResponsibleID: user.ID,
		}
		_, err := repo.Create(&p)

		assert.Error(t, err)
	})
}
