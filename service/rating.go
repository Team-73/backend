package service

import (
	"math"

	"github.com/Team-73/backend/domain/contract"
	"github.com/Team-73/backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type ratingService struct {
	svc *Service
}

//newRatingService return a new instance of the service
func newRatingService(svc *Service) contract.RatingService {
	return &ratingService{
		svc: svc,
	}
}

func (s *ratingService) GetCompanyUserRating(companyID, userID int64) (*entity.Rating, *resterrors.RestErr) {

	ratingResult, err := s.svc.db.Rating().GetCompanyUserRating(companyID, userID)
	if err != nil {
		return nil, err
	}

	return ratingResult, nil
}

func (s *ratingService) UpdateRating(rating entity.Rating) (*entity.Rating, *resterrors.RestErr) {

	rating.TotalRating = s.calculateTotalRating(rating)

	ratingResult, err := s.GetCompanyUserRating(rating.CompanyID, rating.UserID)
	if err != nil {
		if err.Message != "Error 0005: No records find with the parameters" {
			return nil, err
		}
	}

	if ratingResult != nil {
		updatedRating, err := s.svc.db.Rating().UpdateRating(rating)
		if err != nil {
			return nil, err
		}

		return updatedRating, nil
	}

	createdRating, err := s.svc.db.Rating().CreateRating(rating)
	if err != nil {
		return nil, err
	}

	return createdRating, nil
}

func (s *ratingService) calculateTotalRating(rating entity.Rating) float64 {
	var totalRatings int64 = 0
	var totalStars int64 = 0
	var totalRatingValue float64 = 0

	if rating.CustomerService > 0 {
		totalRatings++
		totalStars += rating.CustomerService
	}
	if rating.CompanyClean > 0 {
		totalRatings++
		totalStars += rating.CompanyClean
	}
	if rating.IceBeer > 0 {
		totalRatings++
		totalStars += rating.IceBeer
	}
	if rating.GoodFood > 0 {
		totalRatings++
		totalStars += rating.GoodFood
	}
	if rating.WouldGoBack > 0 {
		totalRatings++
		totalStars += rating.WouldGoBack
	}

	totalRatingValue = float64(totalStars) / float64(totalRatings)

	return toFixed(totalRatingValue, 1)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
