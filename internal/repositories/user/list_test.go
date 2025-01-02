package user_test

import (
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepository(t *testing.T) {
	tx := db.Begin()
	repo := user.NewUserRepository(tx)

	defer tx.Rollback()

	t.Run("Should list users filtered correctly", func(t *testing.T) {
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
		tx.Create(&users)
		filter := models.User{
			Age: 30,
		}

		result, err := repo.List(&filter, 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, result[0].Name, "test-name-2")
		assert.Equal(t, result[1].Name, "test-name-3")
	})
}
