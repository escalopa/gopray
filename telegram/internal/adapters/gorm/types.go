package mygorm

import (
	"time"
)

type User struct {
	TelegramID int64 `gorm:"primary_key,auto_increment:false"`

	FirstName string `gorm:"length:255,gorm:default:''"`
	LastName  string `gorm:"length:255,gorm:default:''"`
	Username  string `gorm:"length:255,gorm:default:''"`

	LastMessageID int64  `gorm:"default:0"`
	LangCode      string `gorm:"length:2;default:'en'"`

	IsSubscribed bool `gorm:"default:false"`
	IsAdmin      bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
