package model

import (
	"gorm.io/gorm"
)

type Lootbox struct {
	gorm.Model
	X int `gorm:"required"` // X coordinate of the lootbox
	Y int `gorm:"required"` // Y coordinate of the lootbox

	// Relations
	UnlocksID uint    `gorm:"not null; unique;"`
	Unlocks   Minicon `gorm:"foreignKey:UnlocksID;"`
	RegionID  uint    `gorm:"not null;"`
	Region    Region  `gorm:"foreignKey:RegionID;"`
}

type GeneratedLootbox struct {
	gorm.Model
	IsOpen bool `gorm:"default:false"` // Is the lootbox open?

	// Relations
	UserID    uint    `gorm:"not null; uniqueIndex:user_lootbox"`
	User      User    `gorm:"foreignKey:UserID;"`
	LootboxID uint    `gorm:"not null; uniqueIndex:user_lootbox"`
	Lootbox   Lootbox `gorm:"foreignKey:LootboxID;"`
}
