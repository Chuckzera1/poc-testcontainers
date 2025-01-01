package user_test

import (
	"context"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/testutils"
	"poc-testcontainers/internal/repositories/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepository(t *testing.T) {
	ctx := context.Background()
	db, err := testutils.NewTestDatabase(ctx, &models.User{})
	if err != nil {
		t.Fatalf("Error getting test db \nReason= %s", err.Error())
	}

	gormDB := db.GormDB
	repo := user.NewUserRepository(gormDB)
	t.Run("Should create user correctly", func(t *testing.T) {
		u := models.User{
			Name: "test-name",
			Age:  20,
		}
		result, err := repo.Create(&u)

		var userCreated models.User
		gormDB.Where("name", "test-name").First(&userCreated)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, userCreated.ID)
		assert.Equal(t, "test-name", userCreated.Name)
		assert.Equal(t, 20, userCreated.Age)
	})
}
