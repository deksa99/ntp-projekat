package request

type CreateAppointment struct {
	UserID       uint `json:"UserID"`
	VehicleID    uint `json:"VehicleID"`
	ServiceID    uint `json:"ServiceID"`
	CarServiceID uint `json:"CarServiceID"`
}
