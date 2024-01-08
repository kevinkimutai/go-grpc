package ports

import (
	"context"

	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/domain"
)

type DBPort interface {
	GetPaymentById(ctx context.Context, id uint) (domain.Payment, error)
	SavePayment(ctx context.Context, domain *domain.Payment) (*domain.Payment, error)
}
