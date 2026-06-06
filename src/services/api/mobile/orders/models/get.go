package models

type GetOrderInput struct {
	ID string `path:"id" doc:"Order ID"`
}

type OrderOutput struct {
	Body Order
}
