package entity

import "time"

// Rating entity data
type Rating struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	CompanyID       int64     `json:"company_id"`
	CustomerService int64     `json:"customer_service"`
	CompanyClean    int64     `json:"company_clean"`
	IceBeer         int64     `json:"ice_beer"`
	GoodFood        int64     `json:"good_food"`
	WouldGoBack     int64     `json:"would_go_back"`
	TotalRating     float64   `json:"total_rating"`
	CreatedAt       time.Time `json:"created_at"`
}
