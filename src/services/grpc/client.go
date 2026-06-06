package grpc

import (
	"app/src/app_settings"
	"errors"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ErrContractNotConfigured = errors.New("grpc contract is not configured")

type Client struct {
	conn *ggrpc.ClientConn
}

func NewClient(settings *app_settings.GrpcSettings) (*Client, error) {
	conn, err := ggrpc.NewClient(settings.OrdersAddress, ggrpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Conn() *ggrpc.ClientConn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}
