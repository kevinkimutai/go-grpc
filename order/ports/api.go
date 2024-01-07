package ports

import "github.com/kevinkimutai/go-grpc/order/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
