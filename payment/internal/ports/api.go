package ports

import (
	"context"

	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (*domain.Payment, error)
}
