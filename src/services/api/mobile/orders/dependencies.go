package orders

import (
	"app/src/core/brokers/common"
	"app/src/services/grpc"
)

type Dependencies struct {
	handler *OrderHandler
}

func NewDependencies(broker common.Broker, grpcClient *grpc.Client) *Dependencies {
	orderRepository := NewOrderRepository(broker, grpcClient)
	orderService := NewOrderService(orderRepository)

	return &Dependencies{
		handler: NewOrderHandler(orderService),
	}
}
