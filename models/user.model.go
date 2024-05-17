package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint64    `gorm:"primaryKey;column:user_id;autoIncrement"`
	Username  string    `gorm:"unique; not null;index"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null;index"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "user"
}
