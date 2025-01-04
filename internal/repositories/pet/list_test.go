package pet_test

import (
	"poc-testcontainers/internal/model"
	"poc-testcontainers/internal/repositories/pet"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepository(t *testing.T) {
	tx := db.Begin()
	repo := pet.NewPetRepository(tx)

	defer tx.Rollback()

	t.Run("Should list pet filtered correctly", func(t *testing.T) {
		users := []model.User{
			{

				Name: "test-name",
				Age:  20,
			},
			{

				Name: "test-name-2",
				Age:  30,
			},
		}
		tx.Create(&users)

		pets := []model.Pet{
			{
				Name:              "test-pet-name",
				Age:               1,
				UserResponsibleID: users[0].ID,
			},
			{
				Name:              "test-pet-name-2",
				Age:               4,
				UserResponsibleID: users[0].ID,
			},
			{
				Name:              "test-pet-name-3",
				Age:               10,
				UserResponsibleID: users[1].ID,
			},
			{
				Name:              "test-pet-name-4",
				Age:               3,
				UserResponsibleID: users[0].ID,
			},
		}
		tx.Create(&pets)

		filter := model.Pet{
			UserResponsibleID: users[0].ID,
		}
		result, err := repo.List(&filter, 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))

		assert.Equal(t, "test-pet-name", result[0].Name)
		assert.Equal(t, "test-name", result[0].UserResponsible.Name)

		assert.Equal(t, "test-pet-name-2", result[1].Name)
		assert.Equal(t, "test-name", result[1].UserResponsible.Name)

		assert.Equal(t, "test-pet-name-4", result[2].Name)
		assert.Equal(t, "test-name", result[2].UserResponsible.Name)
	})
}
