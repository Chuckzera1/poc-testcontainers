package user_test

import (
	"poc-testcontainers/internal/model"
	"poc-testcontainers/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletetRepository(t *testing.T) {
	t.Run("Should delete user correctly", func(t *testing.T) {
		tx := db.Begin()
		repo := user.NewUserRepository(tx)

		defer tx.Rollback()
		users := []model.User{
			{

				Name: "test-name",
				Age:  20,
			},
			{

				Name: "test-name-2",
				Age:  30,
			},
			{

				Name: "test-name-3",
				Age:  30,
			},
			{

				Name: "test-name-4",
				Age:  40,
			},
		}
		tx.Create(&users)

		err := repo.Delete(users[1].ID)

		var result []model.User
		tx.Unscoped().Order("name ASC").Find(&result)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 4, len(result))
		assert.NotZero(t, result[1].DeletedAt)
		assert.Zero(t, result[0].DeletedAt)
		assert.Zero(t, result[2].DeletedAt)
		assert.Zero(t, result[3].DeletedAt)
	})

	t.Run("Should delete user with pet correctly", func(t *testing.T) {
		tx := db.Begin()
		repo := user.NewUserRepository(tx)

		defer tx.Rollback()
		users := []model.User{
			{

				Name: "test-name",
				Age:  20,
			},
			{

				Name: "test-name-2",
				Age:  30,
			},
			{

				Name: "test-name-3",
				Age:  30,
			},
			{

				Name: "test-name-4",
				Age:  40,
			},
		}
		tx.Create(&users)

		pets := []model.Pet{
			{

				Name:              "test-name",
				Age:               20,
				UserResponsibleID: users[0].ID,
			},
			{

				Name:              "test-name-1",
				Age:               30,
				UserResponsibleID: users[0].ID,
			},
			{

				Name:              "test-name-2",
				Age:               30,
				UserResponsibleID: users[1].ID,
			},
		}
		tx.Create(&pets)

		err := repo.Delete(users[0].ID)

		var resultUsers []model.User
		var resultPets []model.Pet

		tx.Unscoped().Order("name ASC").Find(&resultUsers)
		tx.Unscoped().Order("name ASC").Find(&resultPets)

		assert.NoError(t, err)
		assert.NotNil(t, resultUsers)
		assert.Equal(t, 4, len(resultUsers))

		assert.NotZero(t, resultUsers[0].DeletedAt)
		assert.Zero(t, resultUsers[1].DeletedAt)
		assert.Zero(t, resultUsers[2].DeletedAt)
		assert.Zero(t, resultUsers[3].DeletedAt)

		assert.NotNil(t, pets)
		assert.Equal(t, 3, len(resultPets))

		assert.NotZero(t, resultPets[0].DeletedAt)
		assert.NotZero(t, resultPets[1].DeletedAt)
		assert.Zero(t, resultPets[2].DeletedAt)
	})
}
