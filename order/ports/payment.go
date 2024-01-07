package ports

import "e-commerce-order/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
