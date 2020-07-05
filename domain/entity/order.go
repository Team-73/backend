package entity

import "time"

// Order entity
type Order struct {
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
	CompanyID   int64          `json:"company_id"`
	AcceptTip   bool           `json:"accept_tip"`
	TotalTip    float64        `json:"total_tip"`
	TotalPrice  float64        `json:"total_price"`
	Observation string         `json:"observation"`
	CreatedAt   time.Time      `json:"created_at"`
	Products    []OrderProduct `json:"products"`
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
	OrderID     int64     `json:"order_id"`
	CompanyID   int64     `json:"company_id"`
	CompanyName string    `json:"company_name"`
	TotalPrice  float64   `json:"total_price"`
	TotalRating float64   `json:"total_rating"`
	TotalItems  int64     `json:"total_items"`
	CreatedAt   time.Time `json:"created_at"`
}

// OrderDetail entity
type OrderDetail struct {
	CompanyID      int64            `json:"company_id"`
	CompanyName    string           `json:"company_name"`
	ProductsDetail []ProductsDetail `json:"products_detail"`
	SubTotal       float64          `json:"sub_total"`
	TotalTip       float64          `json:"total_tip"`
	TotalPrice     float64          `json:"total_price"`
	CreatedAt      time.Time        `json:"created_at"`
}

// ProductsDetail entity
type ProductsDetail struct {
	Quantity          int64   `json:"quantity"`
	ProductName       string  `json:"product_name"`
	TotalProductPrice float64 `json:"total_product_price"`
}
