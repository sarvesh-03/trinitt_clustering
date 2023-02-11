package models

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model

	Value string `gorm:"required"`
	Row   uint   `gorm:"required; uniqueIndex:idx_entity_parameter_row"`

	// Relations
	EntityID    uint      `gorm:"uniqueIndex:idx_entity_parameter_row"`
	ParameterID uint      `gorm:"uniqueIndex:idx_entity_parameter_row"`
	Parameter   Parameter `gorm:"foreignKey:ParameterID"`
}
