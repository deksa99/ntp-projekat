package api

import (
	"time"
)

type AppointmentInfo struct {
	Id             uint      `json:"Id"`
	UserId         uint      `json:"UserId"`
	FirstName      string    `json:"FirstName"`
	LastName       string    `json:"LastName"`
	ServiceId      uint      `json:"ServiceId"`
	ServiceName    string    `json:"ServiceName"`
	CarServiceId   uint      `json:"CarServiceId"`
	CarServiceName string    `json:"CarServiceName"`
	Street         string    `json:"Street"`
	StartTime      time.Time `json:"StartTime"`
	FinishTime     time.Time `json:"FinishTime"`
	Status         string    `json:"Status"`
}
