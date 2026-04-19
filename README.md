# goTaskForge

A distributed task queue and processing system built with Go. goTaskForge provides a REST API for submitting and tracking asynchronous tasks, backed by PostgreSQL for persistence and Redis for messaging, with a separate worker service for background task processing.

## Features

- Submit tasks via REST API and track their lifecycle (`pending` → `in_progress` → `completed` / `failed`)
- UUID-based task identification
- Retry-count tracking for failed tasks
- Auto-migrated PostgreSQL schema via GORM
- Microservices architecture: decoupled API and worker services
- Fully containerized with Docker Compose

## Architecture

```
HTTP Client
    │
    ▼
┌─────────────┐        ┌─────────────┐
│  API Service │        │   Worker    │
│  (Gin/HTTP)  │        │  Service    │
│             │        │             │
│  Routes     │        │  (planned)  │
│  Handlers   │        └─────────────┘
│  Services   │
└──────┬──────┘
       │
       ▼
┌──────────────┐    ┌──────────────┐
│  PostgreSQL  │    │    Redis     │
│  (Tasks DB)  │    │  (Messaging) │
└──────────────┘    └──────────────┘
```

**Layer responsibilities:**

| Layer | File | Responsibility |
|---|---|---|
| Routes | [api/internal/route/routes.go](api/internal/route/routes.go) | URL mapping |
| Handlers | [api/internal/handler/task_handler.go](api/internal/handler/task_handler.go) | HTTP request/response |
| Services | [api/internal/service/task_service.go](api/internal/service/task_service.go) | Business logic |
| Database | [api/internal/database/postgres.go](api/internal/database/postgres.go) | Persistence |

## Project Structure

```
goTaskForge/
├── api/                        # REST API service
│   ├── cmd/server/main.go      # Entry point
│   └── internal/
│       ├── config/             # Environment config loading
│       ├── database/           # PostgreSQL connection & auto-migration
│       ├── dto/                # Request/response contracts
│       ├── handler/            # HTTP handlers
│       ├── middleware/         # Middleware (extensible)
│       ├── route/              # Route definitions
│       └── service/            # Business logic
├── shared/                     # Shared models across services
│   └── model/task.go           # Task model definition
├── worker/                     # Background worker service (WIP)
└── docker-compose.yml          # Container orchestration
```

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/install/)
- **Or** Go 1.22+ and a running PostgreSQL instance for local development

## Quick Start

### Docker Compose (Recommended)

```bash
git clone https://github.com/Sidi1901/goTaskForge.git
cd goTaskForge

docker-compose up -d
```

The API will be available at `http://localhost:8080`.

```bash
# View logs
docker-compose logs -f api

# Stop all services
docker-compose down
```

### Local Development

```bash
cd api
go mod download

export DBHOST=localhost
export DBUSER=postgres
export DBPASSWORD=postgres
export DBNAME=taskforge
export DBPORT=5432
export DBSSLMode=disable
export SERVERPORT=8080

go run ./cmd/server/main.go
```

## Environment Variables

| Variable | Description | Example |
|---|---|---|
| `DBHOST` | PostgreSQL host | `localhost` |
| `DBUSER` | PostgreSQL username | `postgres` |
| `DBPASSWORD` | PostgreSQL password | `postgres` |
| `DBNAME` | Database name | `taskforge` |
| `DBPORT` | PostgreSQL port | `5432` |
| `DBSSLMode` | PostgreSQL SSL mode | `disable` |
| `SERVERPORT` | API server port | `8080` |

## API Reference

Base path: `/api/v1`

### Create a Task

```
POST /api/v1/tasks
```

**Request body:**
```json
{
  "payload": "your-task-data"
}
```

**Response** `201 Created`:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "pending",
  "payload": "your-task-data",
  "result": "",
  "retry_count": 0,
  "created_at": "2026-04-20T10:00:00Z",
  "updated_at": "2026-04-20T10:00:00Z"
}
```

### Get a Task

```
GET /api/v1/tasks/:id
```

**Response** `200 OK`:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "status": "pending",
  "payload": "your-task-data",
  "result": "",
  "retry_count": 0,
  "created_at": "2026-04-20T10:00:00Z",
  "updated_at": "2026-04-20T10:00:00Z"
}
```

**Response** `404 Not Found` if the task ID does not exist.

### Task Status Values

| Status | Description |
|---|---|
| `pending` | Task submitted, awaiting processing |
| `in_progress` | Task is actively being processed |
| `completed` | Task finished successfully |
| `failed` | Task encountered an error |

## Example Usage

```bash
# Create a task
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"payload": "process-image-12345"}'

# Get task status
curl http://localhost:8080/api/v1/tasks/550e8400-e29b-41d4-a716-446655440000
```

## Database Schema

The schema is automatically migrated on startup via GORM:

```sql
CREATE TABLE tasks (
    id          TEXT PRIMARY KEY,
    status      VARCHAR(20),
    payload     TEXT,
    result      TEXT,
    retry_count INT DEFAULT 0,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);
```

## Tech Stack

| Component | Technology |
|---|---|
| Language | Go 1.22+ |
| HTTP Framework | [Gin](https://github.com/gin-gonic/gin) v1.12 |
| ORM | [GORM](https://gorm.io) v1.31 |
| Database | PostgreSQL 15 |
| Cache / Messaging | Redis 7 |
| Containerization | Docker & Docker Compose |
| ID Generation | [google/uuid](https://github.com/google/uuid) |
| Config | [caarlos0/env](https://github.com/caarlos0/env) |

## Roadmap

- [ ] Worker service implementation
- [ ] Redis-based task queue integration
- [ ] Authentication middleware
- [ ] Request logging middleware
- [ ] Retry mechanism in worker
- [ ] OpenAPI / Swagger documentation
- [ ] Unit and integration tests

## License

MIT
