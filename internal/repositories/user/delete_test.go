package user_test

import (
	"context"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/testutils"
	"poc-testcontainers/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeletetRepository(t *testing.T) {
	ctx := context.Background()
	db, err := testutils.NewTestDatabase(ctx, &models.User{})
	if err != nil {
		t.Fatalf("Error getting test db \nReason= %s", err.Error())
	}

	gormDB := db.GormDB
	repo := user.NewUserRepository(gormDB)
	gormDB.Raw("DELETE FROM users")
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
		gormDB.Create(&users)

		err := repo.Delete(users[1].ID)

		var result []models.User
		gormDB.Find(&result)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
	})
}
