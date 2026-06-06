package services

import (
	commonmodels "app/src/services/api/common/models"
	"app/src/services/api/common/utils"
	"app/src/services/api/mobile/orders/models"
	"app/src/services/api/mobile/orders/repositories"
	"context"

	"github.com/google/uuid"
)

type UpdateOrderService struct {
	repositories *repositories.OrderRepositories
}

func NewUpdateOrderService(repositories *repositories.OrderRepositories) *UpdateOrderService {
	return &UpdateOrderService{repositories: repositories}
}

func (s *UpdateOrderService) Update(ctx context.Context, orderID string, req models.UpdateOrderRequest) (commonmodels.AcceptedResponse, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return commonmodels.AcceptedResponse{}, ErrInvalidOrderID
	}

	notificationID := uuid.NewString()
	notification := utils.NewNotification(notificationID, string(OrderNotificationUpdate), orderID, req)
	if err := s.repositories.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}
