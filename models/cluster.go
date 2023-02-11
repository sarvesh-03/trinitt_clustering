package models

import (
	"gorm.io/gorm"
)

type Cluster struct {
	gorm.Model

	Color string `gorm:"type:varchar(255);required"`

	// Relations
	EntityID uint   `gorm:"required"`
	Entity   Entity `gorm:"foreignKey:EntityID"`
}
