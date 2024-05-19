package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentManagementArcScript struct {
	ContentManagementArcScriptId uint64 `gorm:"primaryKey;autoIncrement;not null;column:content_management_arc_script_id;"`
	ContentManagementArcId       uint64
	ArcScript                    string
	CreatedAt                    time.Time `gorm:"not null"`
	UpdatedAt                    time.Time `gorm:"not null"`
	DeletedAt                    gorm.DeletedAt
}

func (ContentManagementArcScript) TableName() string {
	return "content_management_arc_script"
}
