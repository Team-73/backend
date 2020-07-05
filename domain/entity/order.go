package entity

import "time"

// Order entity
type Order struct {
	ID         int64          `json:"id"`
	UserID     int64          `json:"user_id"`
	CompanyID  int64          `json:"company_id"`
	Rating     float32        `json:"rating"`
	AcceptTip  bool           `json:"accept_tip"`
	TotalTip   float64        `json:"total_tip"`
	TotalPrice float64        `json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
	Products   []OrderProduct `json:"products"`
}

// OrderProduct entity
type OrderProduct struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

// OrdersByUserID entity
type OrdersByUserID struct {
}
