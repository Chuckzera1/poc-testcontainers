package user_test

import (
	"poc-testcontainers/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	cleanUpUserDB(t)
	t.Run("Should create user correctly", func(t *testing.T) {
		u := models.User{
			Name: "test-name",
			Age:  20,
		}
		result, err := repo.Create(&u)

		var userCreated models.User
		db.Where("name", "test-name").First(&userCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, userCreated.ID)
		assert.Equal(t, "test-name", userCreated.Name)
		assert.Equal(t, 20, userCreated.Age)
	})
}
