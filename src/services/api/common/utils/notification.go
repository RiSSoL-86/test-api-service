package utils

import (
	"app/src/core/brokers/common"
	"app/src/services/api/common/models"
	"time"
)

func Accepted(orderID string) models.AcceptedResponse {
	return models.AcceptedResponse{ID: orderID, Status: "accepted"}
}

func NewNotification(id string, notificationType string, entityID string, payload any) common.Notification {
	return common.Notification{
		ID:        id,
		Type:      notificationType,
		EntityID:  entityID,
		Payload:   payload,
		CreatedAt: time.Now().UTC(),
	}
}
