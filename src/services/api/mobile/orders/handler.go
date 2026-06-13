package orders

import (
	commonmodels "app/src/services/api/common/models"
	"app/src/services/api/mobile/orders/models"
	"context"
	"errors"
	"log"

	"github.com/danielgtaylor/huma/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	service *OrderService
}

func NewOrderHandler(service *OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Get(ctx context.Context, input *models.GetOrderInput) (*models.OrderOutput, error) {
	res, err := h.service.Get(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrInvalidOrderID) {
			return nil, huma.Error400BadRequest("invalid order id")
		}
		switch status.Code(err) {
		case codes.NotFound:
			return nil, huma.Error404NotFound("order not found")
		case codes.InvalidArgument:
			return nil, huma.Error400BadRequest("invalid order id")
		}
		log.Printf("Get order error: %v", err)
		return nil, huma.Error503ServiceUnavailable("failed to get order")
	}

	return &models.OrderOutput{Body: res}, nil
}

func (h *OrderHandler) Create(ctx context.Context, input *models.CreateOrderInput) (*commonmodels.AcceptedOutput, error) {
	res, err := h.service.Create(ctx, input.Body)
	if err != nil {
		log.Printf("Create order error: %v", err)
		return nil, huma.Error503ServiceUnavailable("failed to enqueue order creation")
	}

	return &commonmodels.AcceptedOutput{Body: res}, nil
}

func (h *OrderHandler) Update(ctx context.Context, input *models.UpdateOrderInput) (*commonmodels.AcceptedOutput, error) {
	res, err := h.service.Update(ctx, input.ID, input.Body)
	if err != nil {
		if errors.Is(err, ErrInvalidOrderID) {
			return nil, huma.Error400BadRequest("invalid order id")
		}
		log.Printf("Update order error: %v", err)
		return nil, huma.Error503ServiceUnavailable("failed to enqueue order update")
	}

	return &commonmodels.AcceptedOutput{Body: res}, nil
}

func (h *OrderHandler) Delete(ctx context.Context, input *models.DeleteOrderInput) (*commonmodels.AcceptedOutput, error) {
	res, err := h.service.Delete(ctx, input.ID)
	if err != nil {
		if errors.Is(err, ErrInvalidOrderID) {
			return nil, huma.Error400BadRequest("invalid order id")
		}
		log.Printf("Delete order error: %v", err)
		return nil, huma.Error503ServiceUnavailable("failed to enqueue order deletion")
	}

	return &commonmodels.AcceptedOutput{Body: res}, nil
}
