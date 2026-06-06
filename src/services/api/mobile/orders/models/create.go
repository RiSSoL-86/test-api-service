package models

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

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
	if !phonePattern.MatchString(strings.TrimSpace(r.CustomerPhone)) {
		errs = append(errs, fieldError("body.customer_phone", "must be a valid phone number", r.CustomerPhone))
	}
	if !isPriority(r.Priority) {
		errs = append(errs, fieldError("body.priority", "must be one of low, normal, high", r.Priority))
	}
	return errs
}

type CreateOrderInput struct {
	Body CreateOrderRequest
}
