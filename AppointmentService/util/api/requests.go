package api

import (
	"encoding/json"
	"errors"
	rr "github.com/hlts2/round-robin"
	"net/http"
	"net/url"
	"strconv"
)

var userServiceBasePath, _ = rr.New(
	&url.URL{Host: "http://localhost:8081"},
)

var vehicleServiceBasePath, _ = rr.New(
	&url.URL{Host: "http://localhost:8082"},
)

var carServiceServiceBasePath, _ = rr.New(
	&url.URL{Host: "http://localhost:8083"},
)

func getUserDetails(id uint) (UserInfo, error) {
	response, err := http.Get(userServiceBasePath.Next().Host + "/api/users/" + strconv.FormatUint(uint64(id), 10))

	if err != nil {
		return UserInfo{}, err
	}
	if response.StatusCode != 200 {
		return UserInfo{}, errors.New("user not found")
	}

	var user UserInfo
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return UserInfo{}, err
	} else {
		return user, nil
	}
}

func getVehicleDetails(id uint) (VehicleInfo, error) {
	response, err := http.Get(vehicleServiceBasePath.Next().Host + "/api/vehicles/" + strconv.FormatUint(uint64(id), 10))

	if err != nil {
		return VehicleInfo{}, err
	}
	if response.StatusCode != 200 {
		return VehicleInfo{}, errors.New("vehicle not found")
	}

	var vehicle VehicleInfo
	err = json.NewDecoder(response.Body).Decode(&vehicle)
	if err != nil {
		return VehicleInfo{}, err
	} else {
		return vehicle, nil
	}
}

func getCarService(id uint) (CarServiceInfo, error) {
	response, err := http.Get(carServiceServiceBasePath.Next().Host + "/api/car-services/" + strconv.FormatUint(uint64(id), 10))

	if err != nil {
		return CarServiceInfo{}, err
	}
	if response.StatusCode != 200 {
		return CarServiceInfo{}, errors.New("vehicle not found")
	}

	var carService CarServiceInfo
	err = json.NewDecoder(response.Body).Decode(&carService)
	if err != nil {
		return CarServiceInfo{}, err
	} else {
		return carService, nil
	}
}
