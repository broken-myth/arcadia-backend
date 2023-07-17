package model

import (
	"gorm.io/gorm"
)

type Minicon struct {
	gorm.Model
	Name        string `gorm:"not null;"`
	BaseHealth  uint   `gorm:"not null;"`
	BaseAttack  uint   `gorm:"not null;"`
	Description string `gorm:"not null;"`
	ImageLink   string `gorm:"not null;"`

	//Relations
	TypeID uint        `gorm:"not null;"`
	Type   MiniconType `gorm:"foreignKey:TypeID"`
}
