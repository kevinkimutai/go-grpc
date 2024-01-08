package api

import (
	"context"

	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/domain"
	"github.com/kevinkimutai/go-grpc/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (*domain.Payment, error) {
	result, err := a.db.SavePayment(ctx, &payment)
	if err != nil {
		return nil, err
	}
	return result, err
}
