package grpc

import (
	"context"
	"fmt"

	"github.com/kevinkimutai/go-grpc/payment/internal/application/core/domain"
	"github.com/kevinkimutai/go-grpc/payment/internal/proto/golang/payment"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	log.WithContext(ctx).Info("Creating payment...")
	newPayment := domain.NewPayment(uint(request.UserId), uint(request.OrderId), request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: int64(result.ID)}, nil
}
