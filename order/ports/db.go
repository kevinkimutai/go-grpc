package ports

import "github.com/kevinkimutai/go-grpc/order/application/core/domain"

type DBPort interface {
	GetOrderById(id string) (domain.Order, error)
	SaveOrder(*domain.Order) error
}
