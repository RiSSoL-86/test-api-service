package utils

import (
	"app/src/core/brokers"
	"app/src/services/api/common/models"
	"encoding/json"
	"time"
)

func Accepted(orderID string) models.AcceptedResponse {
	return models.AcceptedResponse{ID: orderID, Status: "accepted"}
}

func NewNotification(id string, notificationType string, entityID string, payload any) (brokers.Notification, error) {
	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return brokers.Notification{}, err
	}

	return brokers.Notification{
		NotificationID: id,
		Type:           notificationType,
		EntityID:       entityID,
		Payload:        rawPayload,
		CreatedAt:      time.Now().UTC(),
	}, nil
}
