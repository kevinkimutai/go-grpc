package grpc

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/kevinkimutai/go-grpc/order/application/core/domain"
	"github.com/kevinkimutai/go-grpc/order/proto/golang/order"
)

//Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	log.WithContext(ctx).Info("Creating order...")
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    uint(orderItem.Quantity),
		})
	}
	newOrder := domain.NewOrder(uint(request.UserId), orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: uint64(result.ID)}, nil
}
