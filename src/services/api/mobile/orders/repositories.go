package orders

import (
	"app/src/core/brokers"
	orderspb "app/src/core/proto/orders"
	orderscommon "app/src/services/api/mobile/orders/common"
	"app/src/services/api/mobile/orders/models"
	"app/src/services/grpc"
	"context"
	"encoding/json"
	"time"
)

const (
	topic          = "orders"
	publishTimeout = 10 * time.Second
)

type OrderRepository struct {
	broker     brokers.Broker
	grpcClient *grpc.Client
}

func NewOrderRepository(broker brokers.Broker, grpcClient *grpc.Client) *OrderRepository {
	return &OrderRepository{
		broker:     broker,
		grpcClient: grpcClient,
	}
}

func (r *OrderRepository) Get(ctx context.Context, orderID string) (models.Order, error) {
	response, err := r.grpcClient.Orders().Get(ctx, &orderspb.GetRequest{Id: orderID})
	if err != nil {
		return models.Order{}, err
	}

	return orderscommon.ToOrderModel(response.GetOrder()), nil
}

func (r *OrderRepository) SendNotification(ctx context.Context, notification brokers.Notification) error {
	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, publishTimeout)
	defer cancel()

	return r.broker.Producer().Publish(ctx, brokers.Message{
		Topic: topic,
		Key:   []byte(notification.EntityID),
		Value: body,
	})
}
