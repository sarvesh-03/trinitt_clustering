package models

import (
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model

	Name string `gorm:"type:varchar(255);required"`

	// Relations
	CreatedByID uint `gorm:"required"`
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
}
