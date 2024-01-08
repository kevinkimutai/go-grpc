package api

import (
	"context"

	"github.com/kevinkimutai/go-grpc/order/application/core/domain"
	"github.com/kevinkimutai/go-grpc/order/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	err := a.db.SaveOrder(&order)
	if err != nil {
		return domain.Order{}, err
	}

	//Charge Payment
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}
	return order, nil

}
