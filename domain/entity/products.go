package entity

import (
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// Product entity
type Product struct {
	ID                       int64   `json:"id"`
	Name                     string  `json:"name"`
	Description              string  `json:"description"`
	Price                    float64 `json:"price"`
	DiscountPrice            float64 `json:"discount_price"`
	CategoryID               int64   `json:"category_id"`
	MinimumAgeForConsumption int64   `json:"minimum_age_for_consumption"`
	ProductImageURL          string  `json:"product_image_url"`
	TimeForPreparingMinutes  int64   `json:"time_for_preparing_minutes"`
}

// Validate to validate a user data
func (product *Product) Validate() *resterrors.RestErr {

	product.Name = strings.TrimSpace(strings.ToLower(product.Name))
	if product.Name == "" {
		return resterrors.NewBadRequestError("Name is invalid")
	}

	if product.Price == 0 {
		return resterrors.NewBadRequestError("Price is invalid")
	}

	if product.CategoryID == 0 {
		return resterrors.NewBadRequestError("Category id is invalid")
	}

	return nil
}
