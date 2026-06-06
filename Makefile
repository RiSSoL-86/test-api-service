-include src/.env
export KAFKA_BROKER_ADDRESSES
export GRPC_ORDERS_ADDRESS
export APP_KAFKA_BROKER_ADDRESSES
export APP_GRPC_ORDERS_ADDRESS

install-pre-commit:
	pre-commit install

docker-build-bin:
	powershell -NoProfile -ExecutionPolicy Bypass -Command "$$env:CGO_ENABLED='0'; $$env:GOOS='linux'; $$env:GOARCH='amd64'; go build -o bin/app ./src"

deps:
	docker compose -p test-api-service-deps --env-file src/.env -f compose.deps.dev.yml up -d

app: docker-build-bin
	docker compose -p test-api-service-app --env-file src/.env -f compose.dev.yml up -d

up: deps app

down:
	docker compose -p test-api-service-app --env-file src/.env -f compose.dev.yml down
	docker compose -p test-api-service-deps --env-file src/.env -f compose.deps.dev.yml down

check:
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -race -v ./...

build:
	go build -race -o bin/app.exe ./src

run:
	go run -race ./src
