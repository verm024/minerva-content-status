package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentManagement struct {
	gorm.Model
	ContentManagementId int `gorm:"primaryKey;autoIncrement;not null"`
	Title               string
	Status              string
	Description         string
	CreatedAt           time.Time `gorm:"not null"`
	UpdatedAt           time.Time `gorm:"not null"`
	DeletedAt           gorm.DeletedAt
}

func (ContentManagement) TableName() string {
	return "content_management"
}
