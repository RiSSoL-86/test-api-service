package main

import (
	"app/src/app_settings"
	"app/src/core/brokers"
	"app/src/services/api"
	"app/src/services/grpc"
	"log"
	"net/http"
)

func main() {
	configs, err := app_settings.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	broker := brokers.NewKafkaBroker(configs.Brokers.KafkaSettings)
	defer func() {
		if err := broker.Close(); err != nil {
			log.Printf("Failed to close broker: %v", err)
		}
	}()

	grpcClient, err := grpc.NewClient(configs.GrpcSettings)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer func() {
		if err := grpcClient.Close(); err != nil {
			log.Printf("Failed to close gRPC client: %v", err)
		}
	}()

	apiMux := http.NewServeMux()
	dependencies := api.NewDependencies(broker, grpcClient)

	api.InitAPI(apiMux, dependencies, configs.Api.HumaConfig())
	log.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", apiMux); err != nil {
		log.Fatal(err)
	}
}
