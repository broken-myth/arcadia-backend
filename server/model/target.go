package model

import (
	"gorm.io/gorm"
)

type Target struct {
	gorm.Model
	TargetVal int `gorm:"not null;"`
	//Relations
	PerkID uint `gorm:"not null;"`
	Perk   Perk `gorm:"foreignKey:PerkID;"`
}
