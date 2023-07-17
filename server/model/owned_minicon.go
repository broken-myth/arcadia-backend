package model

import (
	"gorm.io/gorm"
)

type OwnedMinicon struct {
	gorm.Model
	Health uint `gorm:"not null;"`
	Attack uint `gorm:"not null;"`
	XP     uint `gorm:"not null;"`
	Level  uint `gorm:"default:1;"`

	// Relations
	OwnerID   uint    `gorm:"not null; uniqueIndex:owner_minicon"`
	Owner     User    `gorm:"foreignKey:OwnerID;"`
	MiniconID uint    `gorm:"not null; uniqueIndex:owner_minicon"`
	Minicon   Minicon `gorm:"foreignKey:MiniconID;"`
}
