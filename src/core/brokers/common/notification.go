package common

import "time"

type Notification struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	EntityID  string    `json:"entity_id"`
	Payload   any       `json:"payload,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
