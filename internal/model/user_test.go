package model_test

import (
	"context"
	"fmt"
	"os"
	"poc-testcontainers/internal/model"
	"poc-testcontainers/internal/repositories/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserBeforeDelete_DeletesPets(t *testing.T) {
	ctx := context.Background()
	testDB, err := testutils.NewTestDatabase(ctx, &model.User{}, &model.Pet{})
	if err != nil {
		fmt.Printf("Error getting test db \nReason= %s", err.Error())
		os.Exit(1)
	}

	db := testDB.GormDB

	db.AutoMigrate(&model.User{}, &model.Pet{})

	user := model.User{
		Name: "John",
		Age:  30,
	}
	db.Create(&user)

	pets := []model.Pet{
		{
			Name:              "Rex",
			Age:               10,
			UserResponsibleID: user.ID,
		},
		{
			Name:              "Bidu",
			Age:               9,
			UserResponsibleID: user.ID,
		},
	}

	db.Create(&pets)

	db.Delete(&user)

	var petsFound []model.Pet
	db.Unscoped().Find(&petsFound)

	for _, pet := range petsFound {
		assert.NotZero(t, pet.DeletedAt)
	}
}
