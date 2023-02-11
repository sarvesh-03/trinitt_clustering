package models

import (
	"gorm.io/gorm"
)

type ParameterType string

const (
	ParameterTypeString ParameterType = "STRING"
	ParameterTypeInt    ParameterType = "INT"
)

type Parameter struct {
	gorm.Model
	KeyName string        `gorm:"type:varchar(255);required"`
	Type    ParameterType `sql:"type:ENUM('STRING', 'INT');required"`

	// Relations
	EntityID uint   `gorm:"required"`
	Entity   Entity `gorm:"foreignKey:EntityID; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
