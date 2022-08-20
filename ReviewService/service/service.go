package service

import (
	"ReviewService/model"
	"ReviewService/repository"
	"ReviewService/response"
	"ReviewService/util/api"
	"ReviewService/util/converter"
	"github.com/pkg/errors"
)

func AddReview(appointmentID uint, userId uint, rating uint, comment string) (response.ReviewInfo, error) {
	appointment, err := api.GetAppointmentDetails(appointmentID)
	if err != nil {
		return response.ReviewInfo{}, err
	}
	if appointment.UserId != userId {
		return response.ReviewInfo{}, errors.New("not your appointment")
	}
	if appointment.Status != "Finished" {
		return response.ReviewInfo{}, errors.New("appointment not finished")
	}
	if rating > 5 || rating < 1 {
		return response.ReviewInfo{}, errors.New("rating range error")
	}
	review := model.Review{
		AppointmentID:  appointment.Id,
		ServiceID:      appointment.ServiceId,
		ServiceName:    appointment.ServiceName,
		CarServiceID:   appointment.CarServiceId,
		CarServiceName: appointment.CarServiceName,
		UserID:         appointment.UserId,
		FirstName:      appointment.FirstName,
		LastName:       appointment.LastName,
		Rating:         rating,
		Comment:        comment,
	}
	newReview, err := repository.SaveReview(review)
	if err != nil {
		return response.ReviewInfo{}, err
	}

	return converter.ReviewToReviewInfo(&newReview), nil
}

func GetReportedReviews() ([]response.ReportInfo, error) {
	reports := repository.GetReports()

	var infos []response.ReportInfo

	for _, r := range reports {
		infos = append(infos, converter.ReportToReportInfo(&r))
	}

	if len(infos) == 0 {
		return []response.ReportInfo{}, nil
	}

	return infos, nil
}

func ReportReview(id uint) (response.ReviewInfo, error) {
	review, err := repository.GetReviewById(id)
	if err != nil {
		return response.ReviewInfo{}, err
	}
	report := model.ReviewReport{
		Review:        review,
		Processed:     false,
		Inappropriate: false,
	}

	_, err = repository.SaveReport(report)
	if err != nil {
		return response.ReviewInfo{}, err
	}

	return converter.ReviewToReviewInfo(&review), nil
}

func ProcessReport(id uint, inappropriate bool) (response.ReportInfo, error) {
	report, err := repository.GetReportById(id)
	if err != nil {
		return response.ReportInfo{}, err
	}
	if report.Processed {
		return response.ReportInfo{}, errors.New("already processed")
	}

	report.Inappropriate = inappropriate
	report.Processed = true

	_, err = repository.SaveReport(report)
	if err != nil {
		return response.ReportInfo{}, err
	}

	return converter.ReportToReportInfo(&report), nil
}
