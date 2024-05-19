package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentManagementArc struct {
	ContentManagementArcId     uint64 `gorm:"primaryKey;autoIncrement;not null;column:content_management_arc_id;"`
	Title                      string `gorm:"not null;index;"`
	Description                string
	IsFinal                    bool
	IsVoiceRecorded            bool
	IsEdited                   bool
	CreatedAt                  time.Time `gorm:"not null"`
	UpdatedAt                  time.Time `gorm:"not null"`
	DeletedAt                  gorm.DeletedAt
	ContentManagementId        uint64                       `gorm:"not null;not null;index"`
	ContentManagementArcScript []ContentManagementArcScript `gorm:"foreignKey:content_management_arc_id;constraint:OnDelete:CASCADE;"`
}

func (ContentManagementArc) TableName() string {
	return "content_management_arc"
}
