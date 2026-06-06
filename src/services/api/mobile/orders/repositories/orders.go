package repositories

import (
	"app/src/core/brokers/common"
	"app/src/services/api/mobile/orders/models"
	"app/src/services/grpc"
	"context"
	"encoding/json"
)

const Topic = "orders.commands"

type OrderRepositories struct {
	broker     common.Broker
	grpcClient *grpc.Client
}

func NewOrderRepositories(broker common.Broker, grpcClient *grpc.Client) *OrderRepositories {
	return &OrderRepositories{
		broker:     broker,
		grpcClient: grpcClient,
	}
}

func (r *OrderRepositories) Get(_ context.Context, _ string) (models.Order, error) {
	_ = r.grpcClient.Conn()
	return models.Order{}, grpc.ErrContractNotConfigured
}

func (r *OrderRepositories) SendNotification(ctx context.Context, notification common.Notification) error {
	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return r.broker.Producer().Publish(ctx, common.Message{
		Topic:   Topic,
		Key:     []byte(notification.EntityID),
		Value:   body,
		Headers: map[string]string{"message_type": notification.Type},
	})
}
