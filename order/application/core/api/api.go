package api

import (
	"e-commerce-order/application/core/domain"
	"e-commerce-order/ports"
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
