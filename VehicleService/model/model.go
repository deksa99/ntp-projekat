package model

import (
	"gorm.io/gorm"
)

type Manufacturer string

const (
	Toyota       Manufacturer = "Toyota"
	Honda                     = "Honda"
	Chevrolet                 = "Chevrolet"
	Ford                      = "Ford"
	MercedesBenz              = "Mercedes-Benz"
	Jeep                      = "Jeep"
	BMW                       = "BMW"
	Porsche                   = "Porsche"
	Subaru                    = "Subaru"
)

type Vehicle struct {
	gorm.Model
	UserID        uint         `gorm:"not null;default:null"`
	Manufacturer  Manufacturer `sql:"type:manufacturer"`
	CarModel      string       `gorm:"default:null;size:256"`
	Color         string       `gorm:"default:null;size:256"`
	LicencePlate  string       `gorm:"not null;default:null;unique"`
	ChassisNumber string       `gorm:"not null;default:null;unique"`
}
