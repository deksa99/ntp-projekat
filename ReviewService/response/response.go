package response

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}

type ReviewInfo struct {
	ID             uint   `json:"ID"`
	AppointmentID  uint   `json:"AppointmentID"`
	ServiceID      uint   `json:"ServiceID"`
	ServiceName    string `json:"ServiceName"`
	CarServiceName string `json:"CarServiceName"`
	CarServiceID   uint   `json:"CarServiceID"`
	UserID         uint   `json:"UserID"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Rating         uint   `json:"Rating"`
	Comment        string `json:"Comment"`
}

type ReportInfo struct {
	ID             uint   `json:"ID"`
	AppointmentID  uint   `json:"AppointmentID"`
	ServiceID      uint   `json:"ServiceID"`
	ServiceName    string `json:"ServiceName"`
	CarServiceName string `json:"CarServiceName"`
	CarServiceID   uint   `json:"CarServiceID"`
	UserID         uint   `json:"UserID"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Rating         uint   `json:"Rating"`
	Comment        string `json:"Comment"`
	Processed      bool   `json:"Processed"`
	Inappropriate  bool   `json:"Inappropriate"`
}
