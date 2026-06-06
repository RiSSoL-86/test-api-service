package app_settings

import (
	"errors"
	"os"
	"strings"
)

type GrpcSettings struct {
	OrdersAddress string
}

func NewGrpcSettings() (*GrpcSettings, error) {
	ordersAddress := strings.TrimSpace(os.Getenv("GRPC_ORDERS_ADDRESS"))
	if ordersAddress == "" {
		return nil, errors.New("GRPC_ORDERS_ADDRESS is not set")
	}

	return &GrpcSettings{OrdersAddress: ordersAddress}, nil
}
