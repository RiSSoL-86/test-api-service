package grpc

import (
	"app/src/app_settings"
	orderspb "app/src/core/proto/orders"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *ggrpc.ClientConn
	orders orderspb.OrdersServiceClient
}

func NewClient(settings *app_settings.GrpcSettings) (*Client, error) {
	conn, err := ggrpc.NewClient(settings.OrdersAddress, ggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		orders: orderspb.NewOrdersServiceClient(conn),
	}, nil
}

func (c *Client) Orders() orderspb.OrdersServiceClient {
	return c.orders
}

func (c *Client) Close() error {
	return c.conn.Close()
}
