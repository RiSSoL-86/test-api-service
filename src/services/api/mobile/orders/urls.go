package orders

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func InitOrders(api huma.API, dependencies *Dependencies) {
	huma.Register(api, huma.Operation{
		OperationID: "get-order",
		Method:      http.MethodGet,
		Path:        "/api/v1/orders/{id}",
		Summary:     "Get an order",
	}, dependencies.handler.Get)

	huma.Register(api, huma.Operation{
		OperationID:   "create-order",
		Method:        http.MethodPost,
		Path:          "/api/v1/orders",
		Summary:       "Create an order",
		DefaultStatus: http.StatusAccepted,
	}, dependencies.handler.Create)

	huma.Register(api, huma.Operation{
		OperationID:   "update-order",
		Method:        http.MethodPatch,
		Path:          "/api/v1/orders/{id}",
		Summary:       "Update an order",
		DefaultStatus: http.StatusAccepted,
	}, dependencies.handler.Update)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-order",
		Method:        http.MethodDelete,
		Path:          "/api/v1/orders/{id}",
		Summary:       "Delete an order",
		DefaultStatus: http.StatusAccepted,
	}, dependencies.handler.Delete)
}
