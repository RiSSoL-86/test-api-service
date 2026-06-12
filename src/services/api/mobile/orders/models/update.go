package models

import "github.com/danielgtaylor/huma/v2"

type UpdateOrderRequest struct {
	Title         *string `json:"title,omitempty" minLength:"3" maxLength:"120" doc:"Short order title"`
	Description   *string `json:"description,omitempty" minLength:"10" maxLength:"1000" doc:"Work details"`
	CustomerName  *string `json:"customer_name,omitempty" minLength:"2" maxLength:"120" doc:"Customer full name"`
	CustomerPhone *string `json:"customer_phone,omitempty" minLength:"9" maxLength:"24" pattern:"^\\+?[0-9][0-9\\s\\-()]*$" doc:"Customer phone number"`
	Address       *string `json:"address,omitempty" minLength:"5" maxLength:"300" doc:"Work address"`
	Priority      *string `json:"priority,omitempty" enum:"low,normal,high" doc:"Order priority"`
	Status        *string `json:"status,omitempty" enum:"new,in_progress,done,cancelled" doc:"Order status"`
}

func (r *UpdateOrderRequest) Resolve(_ huma.Context) []error {
	if r.Title == nil && r.Description == nil && r.CustomerName == nil && r.CustomerPhone == nil &&
		r.Address == nil && r.Priority == nil && r.Status == nil {
		return []error{&huma.ErrorDetail{Location: "body", Message: "must contain at least one field"}}
	}

	return nil
}

type UpdateOrderInput struct {
	ID   string `path:"id" doc:"Order ID"`
	Body UpdateOrderRequest
}
