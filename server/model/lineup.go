package model

import (
	"gorm.io/gorm"
)

type Lineup struct {
	gorm.Model
	PositionNo uint `gorm:"not null; uniqueIndex:pos_min_owner"`

	// Relations
	CreatorID      uint         `gorm:"not null; uniqueIndex:pos_min_owner"`
	Creator        User         `gorm:"foreignKey:CreatorID"`
	OwnedMiniconID uint         `gorm:"not null; uniqueIndex:pos_min_owner"`
	OwnedMinicon   OwnedMinicon `gorm:"foreignKey:OwnedMiniconID;"`
}
