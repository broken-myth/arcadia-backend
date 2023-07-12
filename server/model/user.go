package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// User
	Trophies uint `gorm:"not null;"`
	XP       uint `gorm:"default:0;"`

	// Relations
	UserRegistrationID uint             `gorm:"not null;"`
	UserRegistration   UserRegistration `gorm:"foreignKey:UserRegistrationID;"`
	ArenaID            uint             `gorm:"default:1;"`
	Arena              Arena            `gorm:"foreignKey:ArenaID;"`
	CharacterID        uint             `gorm:"default:1;"`
	Character          Character        `gorm:"foreignKey:CharacterID;"`
}
