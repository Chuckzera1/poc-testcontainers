package user_test

import (
	"poc-testcontainers/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletetRepository(t *testing.T) {
	cleanUpUserDB(t)
	t.Run("Should delete user correctly", func(t *testing.T) {
		users := []models.User{
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
		db.Create(&users)

		err := repo.Delete(users[1].ID)

		var result []models.User
		db.Find(&result)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
	})
}
