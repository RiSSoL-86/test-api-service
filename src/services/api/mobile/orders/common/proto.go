package common

import (
	orderspb "app/src/core/proto/orders"
	"app/src/services/api/mobile/orders/models"
)

func ToOrderModel(order *orderspb.Order) models.Order {
	if order == nil {
		return models.Order{}
	}

	return models.Order{
		ID:            order.GetId(),
		Title:         order.GetTitle(),
		Description:   order.GetDescription(),
		CustomerName:  order.GetCustomerName(),
		CustomerPhone: order.GetCustomerPhone(),
		Address:       order.GetAddress(),
		Priority:      order.GetPriority(),
		Status:        order.GetStatus(),
		CreatedAt:     order.GetCreatedAt().AsTime(),
		UpdatedAt:     order.GetUpdatedAt().AsTime(),
	}
}
