-include src/.env
export KAFKA_BROKER_ADDRESSES
export GRPC_ORDERS_ADDRESS
export APP_CONTAINER_PORT
export APP_KAFKA_BROKER_ADDRESSES
export APP_GRPC_ORDERS_ADDRESS

# Pre-commit
install-pre-commit:
	pre-commit install

# Local development:
deps:
	docker compose -p test-api-service-local -f compose.dev.local.yml up -d

local-run: deps
	go run -race ./src

# Start the full Docker stack:
docker-build-bin: export CGO_ENABLED=0
docker-build-bin: export GOOS=linux
docker-build-bin: export GOARCH=amd64
docker-build-bin:
	go build -o bin/app ./src

run: docker-build-bin
	docker compose -p test-api-service --env-file src/.env -f compose.dev.yml up -d --build

# Stop the Docker stack:
down:
	docker compose -p test-api-service --env-file src/.env -f compose.dev.yml down
	docker compose -p test-api-service-local -f compose.dev.local.yml down

# Lint
check:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -race -v ./...
