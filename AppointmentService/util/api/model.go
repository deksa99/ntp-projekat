package api

type UserInfo struct {
	UserID    uint   `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

type VehicleInfo struct {
	VehicleID     uint   `json:"VehicleID"`
	Manufacturer  string `json:"Manufacturer"`
	CarModel      string `json:"CarModel"`
	LicencePlate  string `json:"LicencePlate"`
	ChassisNumber string `json:"ChassisNumber"`
	UserID        uint   `json:"UserID"`
}

type ServiceInfo struct {
	ServiceID    uint    `json:"ServiceID"`
	ServiceName  string  `json:"Name"`
	ServicePrice float32 `json:"Price"`
	Available    bool    `json:"Available"`
	CarServiceID uint    `json:"CarServiceID"`
}

type CarServiceInfo struct {
	CarServiceID   uint   `json:"CarServiceID"`
	CarServiceName string `json:"Name"`
	Street         string `json:"Street"`
}
