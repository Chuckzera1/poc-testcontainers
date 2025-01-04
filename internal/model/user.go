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

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.Model(&Pet{}).Where("user_responsible_id = ?", u.ID).Delete(&Pet{}).Error; err != nil {
		return err
	}
	return nil
}
