# Orders API

HTTP service for accepting order write requests and publishing them to the broker.

## Setup

Start only Kafka dependencies:

```powershell
make deps
```

Run the application locally against Docker Kafka:

```powershell
make run
```

Start the full Docker stack:

```powershell
make up
```

Stop the Docker stack:

```powershell
make down
```

Run checks:

```powershell
make check
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
