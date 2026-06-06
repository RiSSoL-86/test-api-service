package api

import (
	"app/src/core/brokers/common"
	"app/src/services/api/mobile"
	"app/src/services/grpc"
)

type Dependencies struct {
	Mobile *mobile.Dependencies
}

func NewDependencies(broker common.Broker, grpcClient *grpc.Client) *Dependencies {
	return &Dependencies{
		Mobile: mobile.NewDependencies(broker, grpcClient),
	}
}
