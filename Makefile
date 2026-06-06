-include src/.env
export KAFKA_BROKER_ADDRESSES
export GRPC_ORDERS_ADDRESS
export APP_CONTAINER_PORT
export APP_KAFKA_BROKER_ADDRESSES
export APP_GRPC_ORDERS_ADDRESS

install-pre-commit:
	pre-commit install

deps:
	docker compose -p test-api-service-local -f compose.dev.local.yml up -d

docker-build-bin:
	powershell -NoProfile -ExecutionPolicy Bypass -Command "$$env:CGO_ENABLED='0'; $$env:GOOS='linux'; $$env:GOARCH='amd64'; go build -o bin/app ./src"

up: docker-build-bin
	docker compose -p test-api-service --env-file src/.env -f compose.dev.yml up -d --build

down:
	docker compose -p test-api-service --env-file src/.env -f compose.dev.yml down
	docker compose -p test-api-service-local -f compose.dev.local.yml down

check:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -race -v ./...

build:
	go build -race -o bin/app.exe ./src

run:
	go run -race ./src
