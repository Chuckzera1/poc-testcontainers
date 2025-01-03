package model

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	ID               uint64 `gorm:"primaryKey;autoIncrement;type:bigserial"`
	Name             string
	Age              int
	UserRespnsibleID uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	UserRespnsible   *User `gorm:"foreignKey:UserRespnsibleID"`
}
