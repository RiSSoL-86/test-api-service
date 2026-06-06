package services

import "app/src/services/api/mobile/orders/repositories"

type OrderNotificationType string

const (
	OrderNotificationCreate OrderNotificationType = "order.create"
	OrderNotificationUpdate OrderNotificationType = "order.update"
	OrderNotificationDelete OrderNotificationType = "order.delete"
)

type OrderServices struct {
	GetOrder    *GetOrderService
	CreateOrder *CreateOrderService
	UpdateOrder *UpdateOrderService
	DeleteOrder *DeleteOrderService
}

func NewOrderServices(repositories *repositories.OrderRepositories) *OrderServices {
	return &OrderServices{
		GetOrder:    NewGetOrderService(repositories),
		CreateOrder: NewCreateOrderService(repositories),
		UpdateOrder: NewUpdateOrderService(repositories),
		DeleteOrder: NewDeleteOrderService(repositories),
	}
}
