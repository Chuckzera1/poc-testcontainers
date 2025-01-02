package user_test

import (
	"context"
	"fmt"
	"os"
	"poc-testcontainers/internal/models"
	"poc-testcontainers/internal/repositories/testutils"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()
	testDB, err := testutils.NewTestDatabase(ctx, &models.User{})
	if err != nil {
		fmt.Printf("Error getting test db \nReason= %s", err.Error())
		os.Exit(1)
	}
	db = testDB.GormDB

	m.Run()

	if err := testDB.Cleanup(ctx); err != nil {
		fmt.Printf("failed to clean up test database: %v\n", err)
	}
}
