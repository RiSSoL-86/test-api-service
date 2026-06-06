package main

import (
	"app/src/app_settings"
	"app/src/core/brokers"
	"app/src/services/api"
	"app/src/services/grpc"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const gracefulShutdownTimeout = 10 * time.Second

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

	server := &http.Server{
		Addr:              configs.Api.Address,
		Handler:           apiMux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-shutdownCtx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Failed to gracefully shutdown server: %v", err)
		}
	}()

	log.Printf("Server starting on %s", configs.Api.Address)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
