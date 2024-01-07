package ports

import "github.com/kevinkimutai/go-grpc/order/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
