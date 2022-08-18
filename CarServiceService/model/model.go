package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Latitude  float64 `gorm:"not null;default:null"`
	Longitude float64 `gorm:"not null;default:null"`
	Street    string  `gorm:"default:null;size:256"`
}

type CarService struct {
	gorm.Model
	Name       string `gorm:"not null;default:null;size:256"`
	LocationID uint
	Location   Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Service struct {
	gorm.Model
	Name         string  `gorm:"not null;default:null;size:256"`
	Description  string  `gorm:"not null;default:null;size:1028"`
	Price        float32 `gorm:"not null;default:null"`
	ExpectedTime uint    `gorm:"default:null"`
	Available    bool    `gorm:"not null;default true"`
	CarServiceID uint
	CarService   CarService `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
