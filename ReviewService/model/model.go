package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	AppointmentID  uint   `gorm:"not null;default:null;unique"`
	ServiceID      uint   `gorm:"not null;default:null"`
	ServiceName    string `gorm:"not null;default:null"`
	CarServiceID   uint   `gorm:"not null;default:null"`
	CarServiceName string `gorm:"not null;default:null"`
	UserID         uint   `gorm:"not null;default:null"`
	FirstName      string `gorm:"not null;default:null"`
	LastName       string `gorm:"not null;default:null"`
	Rating         uint   `gorm:"not null;default:null"`
	Comment        string `gorm:"default:null"`
	Inappropriate  bool   `gorm:"not null;default:false"`
}

type ReviewReport struct {
	gorm.Model
	ReviewID  uint   `gorm:"unique"`
	Review    Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Processed bool   `gorm:"not null;default:false"`
}
