package model

import (
	"gorm.io/gorm"
)

type Lootbox struct {
	gorm.Model
	X      int  `gorm:"required"`      // X coordinate of the lootbox
	Y      int  `gorm:"required"`      // Y coordinate of the lootbox
	IsOpen bool `gorm:"default:false"` // Is the lootbox open?

	// Relations
	UserID    uint    `gorm:"not null;"`
	User      User    `gorm:"foreignKey:UserID;"`
	UnlocksID uint    `gorm:"not null;"`
	Unlocks   Minicon `gorm:"foreignKey:UnlocksID;"`
}
