# Orders API

HTTP service for accepting order write requests and publishing them to the broker.

## Setup

Local development:

```powershell
make deps
make run
```

`make deps` starts only Kafka and Kafka UI from `compose.dev.local.yml`.
`make run` starts the app locally with `go run -race ./src`.

Start the full Docker stack:

```powershell
make up
```

`make up` starts app, Kafka and Kafka UI from `compose.dev.yml`.
The app image builds the Go binary inside Docker.

Stop the Docker stack:

```powershell
make down
```

Run checks:

```powershell
make check
```

## Configuration

```text
KAFKA_BROKER_ADDRESSES=localhost:29092
APP_KAFKA_BROKER_ADDRESSES=kafka:9092
APP_HOST_PORT=8080
APP_CONTAINER_PORT=8080
```

`KAFKA_BROKER_ADDRESSES` is used by local `make run`.
`APP_KAFKA_BROKER_ADDRESSES` is used by the app container in the full Docker stack.
`APP_HOST_PORT:APP_CONTAINER_PORT` publishes the app container to the host.

## UI

Swagger UI:

```text
http://localhost:8080/docs
```

OpenAPI:

```text
http://localhost:8080/openapi.json
http://localhost:8080/openapi.yaml
```

Kafka UI:

```text
http://localhost:8081
```

## API

```text
GET    /api/v1/orders/{id}
POST   /api/v1/orders
PATCH  /api/v1/orders/{id}
DELETE /api/v1/orders/{id}
```

Write requests return `202 Accepted` after the broker confirms the message.
