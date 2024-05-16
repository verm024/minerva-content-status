package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID    uint64         `gorm:"primaryKey;column:user_id; autoIncrement"`
	Username  string         `gorm:"not null"`
	Password  string         `gorm:"unique;not null"`
	Email     string         `gorm:"unique;not null"`
	CreatedAt time.Time      `gorm:"unique;not null"`
	UpdatedAt time.Time      `gorm:"unique;not null"`
	DeletedAt gorm.DeletedAt `gorm:"unique;not null"`
}

func (User) TableName() string {
	return "user"
}
