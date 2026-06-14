package api

import (
	"app/src/services/api/mobile"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/swaggest/swgui/v5emb"
)

func InitAPI(mux *http.ServeMux, dependencies *Dependencies, config huma.Config) {
	api := humago.New(mux, config)
	mobile.InitMobile(api, dependencies.Mobile)

	mux.Handle("/docs/", v5emb.New(config.Info.Title, "/openapi.json", "/docs/"))
}
