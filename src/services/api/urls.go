package api

import (
	"app/src/services/api/mobile"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func InitAPI(mux *http.ServeMux, dependencies *Dependencies, config huma.Config) {
	api := humago.New(mux, config)
	mobile.InitMobile(api, dependencies.Mobile)
}
