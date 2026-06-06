package services

import (
	"app/src/services/api/mobile/orders/models"
	"app/src/services/api/mobile/orders/repositories"
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidOrderID = errors.New("invalid order id")

type GetOrderService struct {
	repositories *repositories.OrderRepositories
}

func NewGetOrderService(repositories *repositories.OrderRepositories) *GetOrderService {
	return &GetOrderService{repositories: repositories}
}

func (s *GetOrderService) Get(ctx context.Context, orderID string) (models.Order, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return models.Order{}, ErrInvalidOrderID
	}

	return s.repositories.Get(ctx, orderID)
}
