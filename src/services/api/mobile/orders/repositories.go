package orders

import (
	"app/src/core/brokers/common"
	"app/src/services/api/mobile/orders/models"
	"app/src/services/grpc"
	"context"
	"encoding/json"
	"time"
)

const (
	topic          = "orders.commands"
	publishTimeout = 10 * time.Second
)

type OrderRepository struct {
	broker     common.Broker
	grpcClient *grpc.Client
}

func NewOrderRepository(broker common.Broker, grpcClient *grpc.Client) *OrderRepository {
	return &OrderRepository{
		broker:     broker,
		grpcClient: grpcClient,
	}
}

func (r *OrderRepository) Get(_ context.Context, _ string) (models.Order, error) {
	// TODO: fetch the order over gRPC from the orders service once it exists.
	return models.Order{}, grpc.ErrContractNotConfigured
}

func (r *OrderRepository) SendNotification(ctx context.Context, notification common.Notification) error {
	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, publishTimeout)
	defer cancel()

	return r.broker.Producer().Publish(ctx, common.Message{
		Topic:   topic,
		Key:     []byte(notification.EntityID),
		Value:   body,
		Headers: map[string]string{"message_type": notification.Type},
	})
}
