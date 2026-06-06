package models

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateOrderRequest struct {
	Title         *string `json:"title,omitempty" doc:"Short order title"`
	Description   *string `json:"description,omitempty" doc:"Work details"`
	CustomerName  *string `json:"customer_name,omitempty" doc:"Customer full name"`
	CustomerPhone *string `json:"customer_phone,omitempty" doc:"Customer phone number"`
	Address       *string `json:"address,omitempty" doc:"Work address"`
	Priority      *string `json:"priority,omitempty" doc:"Order priority"`
	Status        *string `json:"status,omitempty" doc:"Order status"`
}

func (r *UpdateOrderRequest) Resolve(_ huma.Context) []error {
	var errs []error
	if r.Title == nil && r.Description == nil && r.CustomerName == nil && r.CustomerPhone == nil &&
		r.Address == nil && r.Priority == nil && r.Status == nil {
		errs = append(errs, fieldError("body", "must contain at least one field", nil))
	}

	errs = appendStringLenError(errs, "body.title", r.Title, 3, 120)
	errs = appendStringLenError(errs, "body.description", r.Description, 10, 1000)
	errs = appendStringLenError(errs, "body.customer_name", r.CustomerName, 2, 120)
	errs = appendStringLenError(errs, "body.address", r.Address, 5, 300)

	if r.CustomerPhone != nil && !phonePattern.MatchString(strings.TrimSpace(*r.CustomerPhone)) {
		errs = append(errs, fieldError("body.customer_phone", "must be a valid phone number", *r.CustomerPhone))
	}
	if r.Priority != nil && !isPriority(*r.Priority) {
		errs = append(errs, fieldError("body.priority", "must be one of low, normal, high", *r.Priority))
	}
	if r.Status != nil && !isStatus(*r.Status) {
		errs = append(errs, fieldError("body.status", "must be one of new, in_progress, done, cancelled", *r.Status))
	}

	return errs
}

type UpdateOrderInput struct {
	ID   string `path:"id" doc:"Order ID"`
	Body UpdateOrderRequest
}
