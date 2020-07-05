package viewmodel

// CompanyRating viewmodel
type CompanyRating struct {
	UserID          int64   `json:"user_id"`
	CompanyID       int64   `json:"company_id"`
	CustomerService int64   `json:"customer_service"`
	CompanyClean    int64   `json:"company_clean"`
	IceBeer         int64   `json:"ice_beer"`
	GoodFood        int64   `json:"good_food"`
	WouldGoBack     int64   `json:"would_go_back"`
	TotalRating     float32 `json:"total_rating"`
}
