package domain

import "time"

type Payment struct {
	ID         uint      `json:"id"`
	CustomerID uint      `json:"customer_id"`
	Status     string    `json:"status"`
	OrderId    uint      `json:"order_id"`
	TotalPrice float32   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewPayment(customerId uint, orderId uint, totalPrice float32) Payment {
	return Payment{
		CreatedAt:  time.Now(),
		Status:     "PENDING",
		CustomerID: customerId,
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}
}
