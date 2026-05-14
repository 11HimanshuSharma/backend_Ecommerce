package models

import "time"


type OrderStatus string

const (
	StatusPending OrderStatus = "pending"
	StatusPaid OrderStatus = "paid"
	StatusFulfilled OrderStatus = "fulfilled"
	StatusCancelled OrderStatus = "cancelled"
)



type Order struct {
	ID  int64 `json:"id"`
	UserID int64 `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Currency string `json:"currency"`
	Status OrderStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ID int64 `json:"id"`
	OrderID int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	UnitPrice float64 `json:"unit_price"`
	Quantity int `json:"quantity"`
}