package request

type AddReview struct {
	AppointmentID uint   `json:"AppointmentID"`
	ServiceID     uint   `json:"ServiceID"`
	CarServiceID  uint   `json:"CarServiceID"`
	UserID        uint   `json:"UserID"`
	Rating        uint   `json:"Rating"`
	Comment       string `json:"Comment"`
}

type ProcessReport struct {
	Inappropriate bool `json:"Inappropriate"`
}
