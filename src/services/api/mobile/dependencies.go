package mobile

import (
	"app/src/core/brokers/common"
	"app/src/services/api/mobile/orders"
	"app/src/services/grpc"
)

type Dependencies struct {
	Orders *orders.Dependencies
}

func NewDependencies(broker common.Broker, grpcClient *grpc.Client) *Dependencies {
	return &Dependencies{
		Orders: orders.NewDependencies(broker, grpcClient),
	}
}
