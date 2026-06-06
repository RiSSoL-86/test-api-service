package app_settings

import "github.com/danielgtaylor/huma/v2"

type ApiSettings struct {
	Title   string
	Version string
}

func NewApiSettings() *ApiSettings {
	return &ApiSettings{
		Title:   "Orders API",
		Version: "1.0.0",
	}
}

func (s *ApiSettings) HumaConfig() huma.Config {
	config := huma.DefaultConfig(s.Title, s.Version)
	config.DocsRenderer = huma.DocsRendererSwaggerUI
	return config
}
