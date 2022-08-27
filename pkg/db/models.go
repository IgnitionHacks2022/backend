package db

import (
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
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID;preload:false"`
	Type   string `gorm:"not null"`
	Name   string `gorm:"not null"`
}
