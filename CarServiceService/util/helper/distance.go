package helper

import (
	"CarServiceService/model"
	"errors"
	"math"
)

func distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	lat1r := lat1 / (180 / math.Pi)
	lon1r := lon1 / (180 / math.Pi)
	lat2r := lat2 / (180 / math.Pi)
	lon2r := lon2 / (180 / math.Pi)

	dlon := lon2r - lon1r
	dlat := lat2r - lat1r

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)

	c := 2 * math.Asin(math.Sqrt(a))

	return c * 6371
}

func Nearest(lat float64, lon float64, all []model.Location) (model.Location, error) {
	if len(all) == 0 {
		return model.Location{}, errors.New("oops")
	}

	min := all[0]
	minDistance := distance(lat, lon, min.Latitude, min.Longitude)

	for _, l := range all {
		d := distance(lat, lon, l.Latitude, l.Longitude)
		if d < minDistance {
			minDistance = d
			min = l
		}
	}

	return min, nil
}
