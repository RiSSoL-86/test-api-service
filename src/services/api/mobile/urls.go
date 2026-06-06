package mobile

import (
	"app/src/services/api/mobile/orders"

	"github.com/danielgtaylor/huma/v2"
)

func InitMobile(api huma.API, dependencies *Dependencies) {
	orders.InitOrders(api, dependencies.Orders)
}
