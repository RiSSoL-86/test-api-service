package models

import "github.com/danielgtaylor/huma/v2"

type CreateOrderRequest struct {
	Title         string `json:"title" minLength:"3" maxLength:"120" doc:"Short order title"`
	Description   string `json:"description" minLength:"10" maxLength:"1000" doc:"Work details"`
	CustomerName  string `json:"customer_name" minLength:"2" maxLength:"120" doc:"Customer full name"`
	CustomerPhone string `json:"customer_phone" minLength:"9" maxLength:"24" doc:"Customer phone number"`
	Address       string `json:"address" minLength:"5" maxLength:"300" doc:"Work address"`
	Priority      string `json:"priority" enum:"low,normal,high" doc:"Order priority"`
}

func (r *CreateOrderRequest) Resolve(_ huma.Context) []error {
	var errs []error
	errs = appendRequiredStringLenError(errs, "body.title", r.Title, titleMinLength, titleMaxLength)
	errs = appendRequiredStringLenError(errs, "body.description", r.Description, descriptionMinLength, descriptionMaxLength)
	errs = appendRequiredStringLenError(errs, "body.customer_name", r.CustomerName, customerNameMinLength, customerNameMaxLength)
	errs = appendRequiredPhoneError(errs, "body.customer_phone", r.CustomerPhone)
	errs = appendRequiredStringLenError(errs, "body.address", r.Address, addressMinLength, addressMaxLength)
	errs = appendRequiredPriorityError(errs, "body.priority", r.Priority)
	return errs
}

type CreateOrderInput struct {
	Body CreateOrderRequest
}
