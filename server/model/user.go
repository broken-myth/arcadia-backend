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
	Username string `gorm:"unique"`
	Trophies int    `gorm:"not null;"`
	XP       uint   `gorm:"default:0;"`

	// Relations
	UserRegistrationID uint             `gorm:"not null;"`
	UserRegistration   UserRegistration `gorm:"foreignKey:UserRegistrationID;"`
	CharacterID        uint             `gorm:"default:1;"`
	Character          Character        `gorm:"foreignKey:CharacterID;"`
}
