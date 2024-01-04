package ports

import "e-commerce-order/application/core/domain"

type DBPort interface {
	GetOrderById(id string) (domain.Order, error)
	SaveOrder(*domain.Order) error
}
