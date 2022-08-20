package service

import (
	"ReviewService/response"
	"github.com/pkg/errors"
)

func AddReview(AppointmentID uint, ServiceID uint, CarServiceID uint, UserID uint, Rating uint, Comment string) (response.ReviewInfo, error) {
	return response.ReviewInfo{}, errors.New("not implemented")
}

func GetReportedReviews() ([]response.ReviewInfo, error) {
	return []response.ReviewInfo{}, errors.New("not implemented")
}

func ReportReview(id uint) (response.ReviewInfo, error) {
	return response.ReviewInfo{}, errors.New("not implemented")
}

func ProcessReport(id uint, inappropriate bool) (response.ReportInfo, error) {
	return response.ReportInfo{}, errors.New("not implemented")
}
