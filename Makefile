-include src/.env
export KAFKA_BROKER_ADDRESSES
export GRPC_ORDERS_ADDRESS
export APP_CONTAINER_PORT

# Pre-commit
install-pre-commit:
	pre-commit install

# gRPC code generation
proto-tools:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto-gen:
	buf generate

# Local development:
deps:
	docker compose -p test-api-service-local --env-file src/.env -f compose.dev.local.yml up -d

local-run: deps
	go run -race ./src

local-down:
	docker compose -p test-api-service-local --env-file src/.env -f compose.dev.local.yml down

# Start the full Docker stack:
docker-build-bin: export CGO_ENABLED=0
docker-build-bin: export GOOS=linux
docker-build-bin: export GOARCH=amd64
docker-build-bin:
	go build -o bin/app ./src

run: docker-build-bin
	docker compose -p test-service --env-file src/.env -f compose.dev.yml up -d --build

down:
	docker compose -p test-service --env-file src/.env -f compose.dev.yml down

# Lint / test
check:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -race -v ./...
