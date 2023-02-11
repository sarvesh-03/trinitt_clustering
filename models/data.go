package models

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model

	Value string `gorm:"required"`

	// Relations
	ParameterID uint      `gorm:"required"`
	Parameter   Parameter `gorm:"foreignKey:ParameterID"`
}
