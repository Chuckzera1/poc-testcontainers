package testutils

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/gorm"
)

type TestDatabase struct {
	Host      string
	Port      string
	User      string
	Password  string
	Database  string
	container testcontainers.Container
	GormDB    *gorm.DB
}

func NewTestDatabase(ctx context.Context, model ...interface{}) (*TestDatabase, error) {
	container, err := startPostgresContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start PostgreSQL container: %w", err)
	}

	dbConfig, err := getContainerConfig(ctx, container)
	if err != nil {
		container.Terminate(ctx)
		return nil, fmt.Errorf("failed to get container config: %w", err)
	}

	gormDB, err := connectWithGorm(dbConfig)
	if err != nil {
		container.Terminate(ctx)
		return nil, fmt.Errorf("failed to connect with GORM: %w", err)
	}

	if len(model) > 0 {
		if err := autoMigratemodel(gormDB, model); err != nil {
			container.Terminate(ctx)
			return nil, fmt.Errorf("failed to run AutoMigrate: %w", err)
		}
	}

	return &TestDatabase{
		Host:      dbConfig.Host,
		Port:      dbConfig.Port,
		User:      dbConfig.User,
		Password:  dbConfig.Password,
		Database:  dbConfig.Database,
		container: container,
		GormDB:    gormDB,
	}, nil
}

func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60 * time.Second),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func autoMigratemodel(db *gorm.DB, model []interface{}) error {
	return db.AutoMigrate(model...)
}

func (db *TestDatabase) Cleanup(ctx context.Context) error {
	return db.container.Terminate(ctx)
}
