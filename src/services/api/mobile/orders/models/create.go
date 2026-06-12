package models

type CreateOrderRequest struct {
	Title         string `json:"title" minLength:"3" maxLength:"120" doc:"Short order title"`
	Description   string `json:"description" minLength:"10" maxLength:"1000" doc:"Work details"`
	CustomerName  string `json:"customer_name" minLength:"2" maxLength:"120" doc:"Customer full name"`
	CustomerPhone string `json:"customer_phone" minLength:"9" maxLength:"24" pattern:"^\\+?[0-9][0-9\\s\\-()]*$" doc:"Customer phone number"`
	Address       string `json:"address" minLength:"5" maxLength:"300" doc:"Work address"`
	Priority      string `json:"priority" enum:"low,normal,high" doc:"Order priority"`
}

type CreateOrderInput struct {
	Body CreateOrderRequest
}
