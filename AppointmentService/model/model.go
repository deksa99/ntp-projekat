package model

import (
	"database/sql/driver"
	"gorm.io/gorm"
	"time"
)

type RequestStatus string

func (m *RequestStatus) Scan(value interface{}) error {
	*m = RequestStatus(value.([]byte))
	return nil
}

func (m RequestStatus) Value() (driver.Value, error) {
	return string(m), nil
}

type AppointmentStatus string

func (m *AppointmentStatus) Scan(value interface{}) error {
	*m = AppointmentStatus(value.([]byte))
	return nil
}

func (m AppointmentStatus) Value() (driver.Value, error) {
	return string(m), nil
}

const (
	Submitted RequestStatus = "Submitted"
	Accepted                = "Accepted"
	Rejected                = "Rejected"
)

const (
	Scheduled AppointmentStatus = "Scheduled"
	Cancelled                   = "Cancelled"
	Finished                    = "Finished"
)

type AppointmentRequest struct {
	gorm.Model
	UserID    uint          `gorm:"not null;default:null"`
	VehicleID uint          `gorm:"not null;default:null"`
	ServiceID uint          `gorm:"not null;default:null"`
	Status    RequestStatus `sql:"type:request_status"`
}

type Appointment struct {
	gorm.Model
	UserID     uint              `gorm:"not null;default:null"`
	VehicleID  uint              `gorm:"not null;default:null"`
	ServiceID  uint              `gorm:"not null;default:null"`
	WorkerID   uint              `gorm:"not null;default:null"`
	StartTime  time.Time         `gorm:"not null;default:null"`
	FinishTime time.Time         `gorm:"default:null"`
	Status     AppointmentStatus `sql:"type:appointment_status"`
}
