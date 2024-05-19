package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentManagement struct {
	ContentManagementId  uint64 `gorm:"primaryKey;autoIncrement;not null;column:content_management_id"`
	Title                string `gorm:"not null;index"`
	Status               string `gorm:"type:varchar(20);check:status IN ('DRAFT','WORKING_ON','WAIT_PUBLISH','PUBLISHED');not null;index;default:DRAFT"`
	Description          string
	CreatedAt            time.Time `gorm:"not null;index"`
	UpdatedAt            time.Time `gorm:"not null"`
	DeletedAt            gorm.DeletedAt
	TiktokLink           string
	YoutubeLink          string
	IgLink               string
	ContentManagementArc []ContentManagementArc `gorm:"foreignKey:content_management_id;constraint:OnDelete:CASCADE;"`
}

func (ContentManagement) TableName() string {
	return "content_management"
}
