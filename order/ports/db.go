package ports

import "github.com/kevinkimutai/go-grpc/order/application/core/domain"

type DBPort interface {
	GetOrderById(id uint) (domain.Order, error)
	SaveOrder(*domain.Order) error
}
