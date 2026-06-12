package orders

import (
	commonmodels "app/src/services/api/common/models"
	"app/src/services/api/common/utils"
	"app/src/services/api/mobile/orders/models"
	"context"
	"errors"

	"github.com/google/uuid"
)

type OrderNotificationType string

const (
	OrderNotificationCreate OrderNotificationType = "order.create"
	OrderNotificationUpdate OrderNotificationType = "order.update"
	OrderNotificationDelete OrderNotificationType = "order.delete"
)

var ErrInvalidOrderID = errors.New("invalid order id")

type OrderService struct {
	repository *OrderRepository
}

func NewOrderService(repo *OrderRepository) *OrderService {
	return &OrderService{repository: repo}
}

func (s *OrderService) Get(ctx context.Context, orderID string) (models.Order, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return models.Order{}, ErrInvalidOrderID
	}

	return s.repository.Get(ctx, orderID)
}

func (s *OrderService) Create(ctx context.Context, req models.CreateOrderRequest) (commonmodels.AcceptedResponse, error) {
	orderID := uuid.NewString()
	notificationID := uuid.NewString()

	notification := utils.NewNotification(notificationID, string(OrderNotificationCreate), orderID, req)
	if err := s.repository.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}

func (s *OrderService) Update(ctx context.Context, orderID string, req models.UpdateOrderRequest) (commonmodels.AcceptedResponse, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return commonmodels.AcceptedResponse{}, ErrInvalidOrderID
	}

	notificationID := uuid.NewString()
	notification := utils.NewNotification(notificationID, string(OrderNotificationUpdate), orderID, req)
	if err := s.repository.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}

func (s *OrderService) Delete(ctx context.Context, orderID string) (commonmodels.AcceptedResponse, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return commonmodels.AcceptedResponse{}, ErrInvalidOrderID
	}

	notificationID := uuid.NewString()
	notification := utils.NewNotification(notificationID, string(OrderNotificationDelete), orderID, nil)
	if err := s.repository.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}
