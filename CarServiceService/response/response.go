package response

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}

type CarServiceInfo struct {
	Id        uint    `json:"Id"`
	Name      string  `json:"Name"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Street    string  `json:"Street"`
}

type ServiceInfo struct {
	Id           uint    `json:"Id"`
	Name         string  `json:"Name"`
	Description  string  `json:"Description"`
	Price        float32 `json:"Price"`
	ExpectedTime uint    `json:"ExpectedTime"`
	Available    bool    `json:"Available"`
	CarServiceID uint    `json:"CarServiceID"`
}
