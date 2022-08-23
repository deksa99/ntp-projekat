package repository

import (
	"ReviewService/model"
	"ReviewService/util/database"
	"errors"
)

func GetReviewById(id uint) (model.Review, error) {
	var review model.Review

	database.Db.First(&review, id)

	if review.ID == 0 {
		return model.Review{}, errors.New("review not found")
	}

	return review, nil
}

func GetReportById(id uint) (model.ReviewReport, error) {
	var report model.ReviewReport

	database.Db.Preload("Review").First(&report, id)

	if report.ID == 0 {
		return model.ReviewReport{}, errors.New("report not found")
	}

	return report, nil
}

func SaveReport(report model.ReviewReport) (model.ReviewReport, error) {
	r := database.Db.Save(&report)

	if r.Error != nil {
		return report, r.Error
	}

	return report, nil
}

func SaveReview(review model.Review) (model.Review, error) {
	r := database.Db.Save(&review)

	if r.Error != nil {
		return review, r.Error
	}

	return review, nil
}

func GetReports() []model.ReviewReport {
	var reports []model.ReviewReport

	database.Db.Preload("Review").Find(&reports)

	return reports
}

func GetReviews(id uint) []model.Review {
	var reviews []model.Review

	database.Db.Where("car_service_id = ?", id).Find(&reviews)

	return reviews
}
