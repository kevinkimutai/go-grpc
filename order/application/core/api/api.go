package api

import (
	"github.com/kevinkimutai/go-grpc/order/application/core/domain"
	"github.com/kevinkimutai/go-grpc/order/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.SaveOrder(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
