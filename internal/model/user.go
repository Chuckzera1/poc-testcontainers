package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;type:bigserial"`
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
