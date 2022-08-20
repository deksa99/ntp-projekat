package request

type AddReview struct {
	AppointmentID uint   `json:"AppointmentID"`
	Rating        uint   `json:"Rating"`
	Comment       string `json:"Comment"`
}

type ProcessReport struct {
	Inappropriate bool `json:"Inappropriate"`
}
