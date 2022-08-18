package request

type Location struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}

type CreateService struct {
	Name         string  `json:"Name"`
	Description  string  `json:"Description"`
	Price        float32 `json:"Price"`
	ExpectedTime uint    `json:"ExpectedTime"`
	CarServiceId uint    `json:"CarServiceId"`
}

type UpdateService struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type CreateCarService struct {
	Name      string  `json:"Name"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Street    string  `json:"Street"`
}
