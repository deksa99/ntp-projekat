package converter

import (
	"ReviewService/model"
	"ReviewService/response"
)

func ReviewToReviewInfo(r *model.Review) response.ReviewInfo {
	return response.ReviewInfo{
		ID:             r.ID,
		AppointmentID:  r.AppointmentID,
		ServiceID:      r.ServiceID,
		ServiceName:    r.ServiceName,
		CarServiceName: r.CarServiceName,
		CarServiceID:   r.CarServiceID,
		UserID:         r.UserID,
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		Rating:         r.Rating,
		Comment:        r.Comment,
	}
}

func ReportToReportInfo(r *model.ReviewReport) response.ReportInfo {
	return response.ReportInfo{
		ID:             r.ID,
		AppointmentID:  r.Review.AppointmentID,
		ServiceID:      r.Review.ServiceID,
		ServiceName:    r.Review.ServiceName,
		CarServiceName: r.Review.CarServiceName,
		CarServiceID:   r.Review.CarServiceID,
		UserID:         r.Review.UserID,
		FirstName:      r.Review.FirstName,
		LastName:       r.Review.LastName,
		Rating:         r.Review.Rating,
		Comment:        r.Review.Comment,
		Processed:      r.Processed,
		Inappropriate:  r.Inappropriate,
	}
}
