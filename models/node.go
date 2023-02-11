package models

import (
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model

	X int `gorm:"required"`
	Y int `gorm:"required"`

	// Relations
	ClusterID uint    `gorm:"required"`
	Cluster   Cluster `gorm:"foreignKey:ClusterID"`
}
