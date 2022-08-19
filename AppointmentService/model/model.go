package model

import (
	"gorm.io/gorm"
	"time"
)

type RequestStatus string

type AppointmentStatus string

const (
	Submitted        RequestStatus = "Submitted"
	CancelledRequest RequestStatus = "Cancelled"
	Accepted         RequestStatus = "Accepted"
	Rejected         RequestStatus = "Rejected"
)

const (
	Scheduled            AppointmentStatus = "Scheduled"
	CancelledAppointment AppointmentStatus = "Cancelled"
	Finished             AppointmentStatus = "Finished"
)

type AppointmentRequest struct {
	gorm.Model
	UserID         uint          `gorm:"not null;default:null"`
	FirstName      string        `gorm:"not null;default:null"`
	LastName       string        `gorm:"not null;default:null"`
	Email          string        `gorm:"not null;default:null"`
	VehicleID      uint          `gorm:"not null;default:null"`
	Manufacturer   string        `gorm:"not null;default:null"`
	CarModel       string        `gorm:"default:null"`
	LicencePlate   string        `gorm:"not null;default:null"`
	ChassisNumber  string        `gorm:"not null;default:null"`
	ServiceID      uint          `gorm:"not null;default:null"`
	ServiceName    string        `gorm:"not null;default:null"`
	ServicePrice   float32       `gorm:"not null;default:null"`
	Available      bool          `gorm:"not null;default:false"`
	CarServiceID   uint          `gorm:"not null;default:null"`
	CarServiceName string        `gorm:"not null;default:null"`
	Street         string        `gorm:"default:null"`
	Status         RequestStatus `sql:"type:request_status"`
	SubmittedAt    time.Time     `gorm:"not null;default:null"`
}

type Appointment struct {
	gorm.Model
	AppointmentRequestID uint
	AppointmentRequest   AppointmentRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	StartTime            time.Time          `gorm:"not null;default:null"`
	FinishTime           time.Time          `gorm:"default:null"`
	Status               AppointmentStatus  `sql:"type:appointment_status"`
}
