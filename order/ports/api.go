package ports

import (
	"context"

	"github.com/kevinkimutai/go-grpc/order/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
}
