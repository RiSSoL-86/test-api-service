package models

import "github.com/danielgtaylor/huma/v2"

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

	errs = appendOptionalStringLenError(errs, "body.title", r.Title, titleMinLength, titleMaxLength)
	errs = appendOptionalStringLenError(errs, "body.description", r.Description, descriptionMinLength, descriptionMaxLength)
	errs = appendOptionalStringLenError(errs, "body.customer_name", r.CustomerName, customerNameMinLength, customerNameMaxLength)
	errs = appendOptionalPhoneError(errs, "body.customer_phone", r.CustomerPhone)
	errs = appendOptionalStringLenError(errs, "body.address", r.Address, addressMinLength, addressMaxLength)
	errs = appendOptionalPriorityError(errs, "body.priority", r.Priority)
	errs = appendOptionalStatusError(errs, "body.status", r.Status)

	return errs
}

type UpdateOrderInput struct {
	ID   string `path:"id" doc:"Order ID"`
	Body UpdateOrderRequest
}
