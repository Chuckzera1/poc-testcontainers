package pet_test

import (
	"context"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/pet"
	"poc-testcontainers/internal/repositories/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListRepository(t *testing.T) {
	ctx := context.Background()
	db, err := testutils.NewTestDatabase(ctx, &models.Pet{})
	if err != nil {
		t.Fatalf("Error getting test db \nReason= %s", err.Error())
	}

	gormDB := db.GormDB
	repo := pet.NewPetRepository(gormDB)
	t.Run("Should list pet filtered correctly", func(t *testing.T) {
		users := []models.User{
			{

				Name: "test-name",
				Age:  20,
			},
			{

				Name: "test-name-2",
				Age:  30,
			},
		}
		gormDB.Create(&users)

		pets := []models.Pet{
			{
				Name:             "test-pet-name",
				Age:              1,
				UserRespnsibleID: users[0].ID,
			},
			{
				Name:             "test-pet-name-2",
				Age:              4,
				UserRespnsibleID: users[0].ID,
			},
			{
				Name:             "test-pet-name-3",
				Age:              10,
				UserRespnsibleID: users[1].ID,
			},
			{
				Name:             "test-pet-name-4",
				Age:              3,
				UserRespnsibleID: users[0].ID,
			},
		}
		gormDB.Create(&pets)

		filter := models.Pet{
			UserRespnsibleID: users[0].ID,
		}
		result, err := repo.List(&filter, 0)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 3, len(result))
		assert.Equal(t, "test-pet-name", result[0].Name)
		assert.Equal(t, "test-pet-name-2", result[1].Name)
		assert.Equal(t, "test-pet-name-4", result[2].Name)
	})
}
