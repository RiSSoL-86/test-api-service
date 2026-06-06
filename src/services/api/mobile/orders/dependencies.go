package orders

import (
	"app/src/core/brokers/common"
	"app/src/services/api/mobile/orders/repositories"
	"app/src/services/api/mobile/orders/services"
	"app/src/services/grpc"
)

type Dependencies struct {
	handler *OrderHandler
}

func NewDependencies(broker common.Broker, grpcClient *grpc.Client) *Dependencies {
	orderRepositories := repositories.NewOrderRepositories(broker, grpcClient)
	orderServices := services.NewOrderServices(orderRepositories)

	return &Dependencies{
		handler: NewOrderHandler(orderServices),
	}
}
