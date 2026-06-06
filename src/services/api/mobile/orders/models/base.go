package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

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

const (
	titleMinLength        = 3
	titleMaxLength        = 120
	descriptionMinLength  = 10
	descriptionMaxLength  = 1000
	customerNameMinLength = 2
	customerNameMaxLength = 120
	customerPhoneMinLen   = 9
	customerPhoneMaxLen   = 24
	addressMinLength      = 5
	addressMaxLength      = 300
)

var phonePattern = regexp.MustCompile(`^\+?[0-9][0-9\s\-()]*$`)

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

func appendRequiredStringLenError(errs []error, location string, value string, minLength int, maxLength int) []error {
	return appendStringLenError(errs, location, &value, minLength, maxLength)
}

func appendOptionalStringLenError(errs []error, location string, value *string, minLength int, maxLength int) []error {
	return appendStringLenError(errs, location, value, minLength, maxLength)
}

func appendOptionalPhoneError(errs []error, location string, value *string) []error {
	if value == nil {
		return errs
	}

	if !phonePattern.MatchString(strings.TrimSpace(*value)) {
		return append(errs, fieldError(location, "must be a valid phone number", *value))
	}

	trimmed := strings.TrimSpace(*value)
	if len(trimmed) < customerPhoneMinLen || len(trimmed) > customerPhoneMaxLen {
		return append(errs, fieldError(
			location,
			fmt.Sprintf("length must be between %d and %d", customerPhoneMinLen, customerPhoneMaxLen),
			*value,
		))
	}

	return errs
}

func appendRequiredPhoneError(errs []error, location string, value string) []error {
	return appendOptionalPhoneError(errs, location, &value)
}

func appendOptionalPriorityError(errs []error, location string, value *string) []error {
	if value == nil {
		return errs
	}

	if !isPriority(*value) {
		return append(errs, fieldError(location, "must be one of low, normal, high", *value))
	}

	return errs
}

func appendRequiredPriorityError(errs []error, location string, value string) []error {
	return appendOptionalPriorityError(errs, location, &value)
}

func appendOptionalStatusError(errs []error, location string, value *string) []error {
	if value == nil {
		return errs
	}

	if !isStatus(*value) {
		return append(errs, fieldError(location, "must be one of new, in_progress, done, cancelled", *value))
	}

	return errs
}

func appendStringLenError(errs []error, location string, value *string, minLength int, maxLength int) []error {
	if value == nil {
		return errs
	}

	trimmed := strings.TrimSpace(*value)
	length := utf8.RuneCountInString(trimmed)
	if length < minLength || length > maxLength {
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
