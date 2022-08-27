package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"not null;unique"`
	Password    string `gorm:"not null"`
	BluetoothID string `gorm:"not null"`
}

type Item struct {
	gorm.Model
	User User      `gorm:"foreignKey:UserID;preload:false"`
	Time time.Time `gorm:"not null"`
	Type string    `gorm:"not null"`
	Name string    `gorm:"not null"`
}