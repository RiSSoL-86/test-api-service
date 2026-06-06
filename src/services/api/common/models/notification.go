package models

type AcceptedResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type AcceptedOutput struct {
	Body AcceptedResponse
}
