package services

import (
	commonmodels "app/src/services/api/common/models"
	"app/src/services/api/common/utils"
	"app/src/services/api/mobile/orders/repositories"
	"context"

	"github.com/google/uuid"
)

type DeleteOrderService struct {
	repositories *repositories.OrderRepositories
}

func NewDeleteOrderService(repositories *repositories.OrderRepositories) *DeleteOrderService {
	return &DeleteOrderService{repositories: repositories}
}

func (s *DeleteOrderService) Delete(ctx context.Context, orderID string) (commonmodels.AcceptedResponse, error) {
	if _, err := uuid.Parse(orderID); err != nil {
		return commonmodels.AcceptedResponse{}, ErrInvalidOrderID
	}

	notificationID := uuid.NewString()
	notification := utils.NewNotification(notificationID, string(OrderNotificationDelete), orderID, nil)
	if err := s.repositories.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}
