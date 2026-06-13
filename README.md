# Orders API

Public HTTP service. It accepts order write requests (POST/PATCH/DELETE) and
publishes them to Kafka, and serves reads (GET) by calling the internal
`test-worker-service` over gRPC.

```
client ──POST/PATCH/DELETE──▶ this API ──▶ Kafka ──▶ worker ──▶ Postgres
client ────────GET─────────▶ this API ──gRPC──▶ worker ──▶ Postgres
```

This service starts Kafka; the worker starts Postgres. Each service is fully
independent with its own `compose.dev.yml`; you bring them up separately
(`make run` in each). Both use host networking (`network_mode: host`), so they
share the host's network namespace and reach each other over `localhost` — no
shared network to create, no inter-service ports to publish. This is the only
service exposed to clients (Swagger on `8080`, Kafka UI on `8081`).

## Setup

Local development:

```powershell
make local-run
```

`make deps` starts only Kafka and Kafka UI from `compose.dev.local.yml`.
`make local-run` starts the app locally with `go run -race ./src`.

Start the Docker stack (run this in each service; order does not matter — the
worker's consumer retries until Kafka is up):

```powershell
make run    # builds the binary, starts this service's compose.dev.yml
```

`make run` here starts this service only (app + Kafka + Kafka UI); run `make run`
in `test-worker-service` for the worker + Postgres, then `make migrate-up` there.
The app image is built from a prebuilt binary (`make docker-build-bin`).

Stop the Docker stack:

```powershell
make down
```

Run checks:

```powershell
make check
```

## gRPC contract

The contract lives in `proto/orders/orders.proto` (no API versioning) and is
shared with `test-worker-service`. Regenerate the Go code after editing the proto:

```powershell
make proto-tools   # installs buf + protoc-gen-go(-grpc) via the Go toolchain (once)
make proto-gen     # regenerates src/proto/**
```

## Configuration

```text
KAFKA_BROKER_ADDRESSES=localhost:9092   # Kafka bootstrap server
GRPC_ORDERS_ADDRESS=localhost:50051     # the worker's gRPC server
APP_CONTAINER_PORT=8080                 # HTTP port the API binds
```

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
`GET` returns `404` when the worker has no such order.
