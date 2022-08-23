package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username     string    `gorm:"not null;default:null;minSize:4;maxSize:128;unique"`
	Password     string    `gorm:"not null;default:null;minSize:6;maxSize:128"`
	Active       bool      `gorm:"not null;default:false"`
	BlockedUntil time.Time `gorm:"default:null"`
}

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;default:null;size:256"`
	LastName  string `gorm:"not null;default:null;size:256"`
	Email     string `gorm:"not null;default:null;unique"`
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type ServiceWorker struct {
	gorm.Model
	FirstName    string `gorm:"not null;default:null;size:256"`
	LastName     string `gorm:"not null;default:null;size:256"`
	Email        string `gorm:"not null;default:null;unique"`
	Main         bool   `gorm:"not null;default:false"`
	CarServiceID uint   `gorm:"not null;default:null"`
	AccountID    uint
	Account      Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Admin struct {
	gorm.Model
	FirstName string `gorm:"not null;default:null;size:256"`
	LastName  string `gorm:"not null;default:null;size:256"`
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
