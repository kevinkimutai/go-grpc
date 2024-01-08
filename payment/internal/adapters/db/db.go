package db

import (
	"context"
	"fmt"

	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID uint
	Status     string
	OrderID    uint
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) GetPaymentById(ctx context.Context, id uint) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.WithContext(ctx).First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         uint(paymentEntity.ID),
		CustomerID: uint(paymentEntity.CustomerID),
		Status:     paymentEntity.Status,
		OrderId:    uint(paymentEntity.OrderID),
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt,
	}
	return payment, res.Error

}
func (a Adapter) SavePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	orderModel := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderID:    uint(payment.OrderId),
		TotalPrice: payment.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		payment.ID = orderModel.ID
	}
	return payment, res.Error

}
