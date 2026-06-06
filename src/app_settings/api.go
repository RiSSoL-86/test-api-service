package app_settings

import (
	"os"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type ApiSettings struct {
	Title   string
	Version string
	Address string
}

func NewApiSettings() *ApiSettings {
	port := strings.TrimSpace(os.Getenv("APP_CONTAINER_PORT"))
	if port == "" {
		port = "8080"
	}

	return &ApiSettings{
		Title:   "Orders API",
		Version: "1.0.0",
		Address: ":" + port,
	}
}

func (s *ApiSettings) HumaConfig() huma.Config {
	config := huma.DefaultConfig(s.Title, s.Version)
	config.DocsRenderer = huma.DocsRendererSwaggerUI
	return config
}
