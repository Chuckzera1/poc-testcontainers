package user_test

import (
	"poc-testcontainers/internal/model"
	"poc-testcontainers/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletetRepository(t *testing.T) {
	tx := db.Begin()
	repo := user.NewUserRepository(tx)

	defer tx.Rollback()

	t.Run("Should delete user correctly", func(t *testing.T) {
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
		tx.Find(&result)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
	})
}
