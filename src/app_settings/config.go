package app_settings

import (
	"app/src/app_settings/brokers"

	"github.com/joho/godotenv"
)

type Config struct {
	Brokers      *brokers.Settings
	GrpcSettings *GrpcSettings
	Api          *ApiSettings
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load("src/.env")

	brokerSettings, err := brokers.NewSettings()
	if err != nil {
		return nil, err
	}

	grpcSettings, err := NewGrpcSettings()
	if err != nil {
		return nil, err
	}

	apiSettings := NewApiSettings()

	return &Config{
		Brokers:      brokerSettings,
		GrpcSettings: grpcSettings,
		Api:          apiSettings,
	}, nil
}
