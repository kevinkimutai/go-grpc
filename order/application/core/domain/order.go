package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint    `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerID uint        `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
}
