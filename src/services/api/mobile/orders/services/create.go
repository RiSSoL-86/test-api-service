package services

import (
	commonmodels "app/src/services/api/common/models"
	"app/src/services/api/common/utils"
	"app/src/services/api/mobile/orders/models"
	"app/src/services/api/mobile/orders/repositories"
	"context"

	"github.com/google/uuid"
)

type CreateOrderService struct {
	repositories *repositories.OrderRepositories
}

func NewCreateOrderService(repositories *repositories.OrderRepositories) *CreateOrderService {
	return &CreateOrderService{repositories: repositories}
}

func (s *CreateOrderService) Create(ctx context.Context, req models.CreateOrderRequest) (commonmodels.AcceptedResponse, error) {
	orderID := uuid.NewString()
	notificationID := uuid.NewString()

	notification := utils.NewNotification(notificationID, string(OrderNotificationCreate), orderID, req)
	if err := s.repositories.SendNotification(ctx, notification); err != nil {
		return commonmodels.AcceptedResponse{}, err
	}

	return utils.Accepted(orderID), nil
}
