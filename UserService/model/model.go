package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username    string    `gorm:"not null;default:null;size:4..128"`
	Password    string    `gorm:"not null;default:null;size:6..128"`
	Active      bool      `gorm:"not null;default:false"`
	BannedUntil time.Time `gorm:"default:null"`
}

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;default:null;size:256"`
	LastName  string `gorm:"not null;default:null;size:256"`
	Email     string `gorm:"not null;default:null"`
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type ServiceWorker struct {
	gorm.Model
	FirstName string `gorm:"not null;default:null;size:256"`
	LastName  string `gorm:"not null;default:null;size:256"`
	Main      bool
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Admin struct {
	gorm.Model
	FirstName string `gorm:"not null;default:null;size:256"`
	LastName  string `gorm:"not null;default:null;size:256"`
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
