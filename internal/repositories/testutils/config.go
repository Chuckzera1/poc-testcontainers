package testutils

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getContainerConfig(ctx context.Context, container testcontainers.Container) (*dbConfig, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	return &dbConfig{
		Host:     host,
		Port:     port.Port(),
		User:     "testuser",
		Password: "testpass",
		Database: "testdb",
	}, nil
}

func connectWithGorm(config *dbConfig) (*gorm.DB, error) {
	dsn := config.DSN()

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (db *dbConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Database,
	)
}
