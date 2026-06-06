package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

const (
	PriorityLow    = "low"
	PriorityNormal = "normal"
	PriorityHigh   = "high"

	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
	StatusCancelled  = "cancelled"
)

var phonePattern = regexp.MustCompile(`^\+?[0-9][0-9\s\-()]{8,20}$`)

type Order struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	Address       string    `json:"address"`
	Priority      string    `json:"priority"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func isPriority(value string) bool {
	switch value {
	case PriorityLow, PriorityNormal, PriorityHigh:
		return true
	default:
		return false
	}
}

func isStatus(value string) bool {
	switch value {
	case StatusNew, StatusInProgress, StatusDone, StatusCancelled:
		return true
	default:
		return false
	}
}

func appendStringLenError(errs []error, location string, value *string, minLength int, maxLength int) []error {
	if value == nil {
		return errs
	}

	trimmed := strings.TrimSpace(*value)
	if len(trimmed) < minLength || len(trimmed) > maxLength {
		return append(errs, fieldError(
			location,
			fmt.Sprintf("length must be between %d and %d", minLength, maxLength),
			*value,
		))
	}

	return errs
}

func fieldError(location string, message string, value any) error {
	return &huma.ErrorDetail{
		Location: location,
		Message:  message,
		Value:    value,
	}
}
