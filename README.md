# 🛍️ Shikposh - Enterprise E-Commerce Backend

> A high-performance, scalable e-commerce backend built with Go, implementing Clean Architecture, DDD, CQRS, and Event-Driven patterns. Designed for production-ready applications with comprehensive monitoring and observability.

[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v3-00ADD8?style=flat-square)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-336791?style=flat-square&logo=postgresql)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-6+-DC382D?style=flat-square&logo=redis)](https://redis.io/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)

### 🔗 Related Projects

- **Frontend Repository**: [shikposh](https://github.com/ali-mahdavi-dev/shikposh) - Modern e-commerce frontend built with Next.js 15, React 19, and TypeScript

---

## ✨ Features

### 🏗️ Architecture & Design Patterns

- 🎯 **Clean Architecture** - Complete separation of concerns with layered architecture
- 🧩 **Domain-Driven Design (DDD)** - Rich domain models with business logic
- 📊 **CQRS** - Command Query Responsibility Segregation for optimized reads/writes
- 🎪 **Event-Driven Architecture** - Asynchronous event processing with message bus
- 🔄 **Repository Pattern** - Abstraction layer for data access
- 💼 **Unit of Work Pattern** - Transaction management and event collection
- 🚌 **Message Bus Pattern** - Centralized command/event routing

### 🚀 Performance & Scalability

- ⚡ **Fiber v3** - Ultra-fast HTTP framework based on FastHTTP
- 🔥 **Concurrent Processing** - Goroutine-based request handling
- 💾 **Connection Pooling** - Optimized database connections
- 🗄️ **Redis Caching** - High-performance caching strategy
- 📡 **Async Event Processing** - Non-blocking event handlers
- 🎯 **Optimized Queries** - Efficient database queries with GORM

### 🔐 Security & Authentication

- 🔑 **JWT Authentication** - Secure token-based authentication
- 🔒 **bcrypt Password Hashing** - Industry-standard password security
- ✅ **Input Validation** - Comprehensive request validation
- 🛡️ **Secure Error Handling** - No sensitive data leakage
- 🔐 **Session Management** - Redis-based session storage

### 📊 Monitoring & Observability

- 📈 **Prometheus** - Metrics collection and monitoring
- 📊 **Grafana** - Beautiful monitoring dashboards
- 📝 **ELK Stack** - Centralized logging (Elasticsearch, Filebeat, Kibana)
- 🔍 **Jaeger** - Distributed tracing with OpenTelemetry
- 📡 **Kafka** - Event streaming for microservices

### 🛠️ Developer Experience

- 📚 **Swagger/OpenAPI** - Interactive API documentation
- 🐳 **Docker & Docker Compose** - Easy development setup
- 🔄 **Database Migrations** - Version-controlled schema management
- 🧪 **Testing Support** - Unit and integration test infrastructure
- 📦 **Modular Design** - Easy to extend and maintain

---

## 🛠️ Tech Stack

<div align="center">

![System Architecture](docs/apwp_aa01.png)

_System Architecture & Technology Stack_

</div>

### Core Framework

| Technology | Version     | Purpose                               |
| ---------- | ----------- | ------------------------------------- |
| **Go**     | 1.25        | High-performance programming language |
| **Fiber**  | v3.0.0-rc.2 | Fast HTTP web framework               |
| **GORM**   | 1.31.0      | Powerful ORM for database operations  |

### Database & Cache

| Technology     | Version | Purpose                        |
| -------------- | ------- | ------------------------------ |
| **PostgreSQL** | 12+     | Primary relational database    |
| **Redis**      | 6+      | Caching and session management |
| **SQLite**     | -       | Development/testing database   |

### Infrastructure & DevOps

| Technology         | Purpose                       |
| ------------------ | ----------------------------- |
| **Docker**         | Containerization              |
| **Docker Compose** | Multi-container orchestration |
| **Prometheus**     | Metrics collection            |
| **Grafana**        | Monitoring dashboards         |
| **ELK Stack**      | Log aggregation and analysis  |
| **Jaeger**         | Distributed tracing           |
| **Kafka**          | Event streaming platform      |

### Libraries & Tools

| Technology                   | Purpose                      |
| ---------------------------- | ---------------------------- |
| **JWT (golang-jwt)**         | Authentication tokens        |
| **Zerolog**                  | Structured logging           |
| **Viper**                    | Configuration management     |
| **Cobra**                    | CLI framework                |
| **Swagger**                  | API documentation            |
| **WebSocket (go-socket.io)** | Real-time communication      |
| **Sarama**                   | Kafka client                 |
| **OpenTelemetry**            | Observability framework      |
| **Jaeger Exporter**          | Distributed tracing exporter |

---

## 🏗️ Architecture

### Clean Architecture Layers

The project follows **Clean Architecture** principles with clear separation of concerns:

```mermaid
graph TD
    A[Presentation Layer<br/>HTTP Handlers, Routes, Middleware] --> B[Application Layer<br/>Command Handlers, Event Handlers]
    B --> C[Domain Layer<br/>Entities, Business Logic, Events]
    C --> D[Infrastructure Layer<br/>Database, Cache, Message Bus, Logging]

    style A fill:#e1f5ff
    style B fill:#fff4e1
    style C fill:#e8f5e9
    style D fill:#fce4ec
```

### Design Patterns

#### 1. Domain-Driven Design (DDD)

- **Rich Domain Models** - Entities with business logic
- **Aggregates** - Product as Aggregate Root
- **Domain Events** - Decoupled event-driven communication
- **Repository Pattern** - Abstracted data access

#### 2. CQRS (Command Query Responsibility Segregation)

Separate read and write operations for optimal performance:

```mermaid
graph TD
    A[HTTP Request] --> B{Command or Query?}
    B -->|Write| C[Command Handler]
    B -->|Read| D[Query Handler]
    C --> E[Domain Logic]
    D --> F[Repository Read]
    E --> G[Repository Write]

    style C fill:#ffcdd2
    style D fill:#c8e6c9
```

#### 3. Event-Driven Architecture

Asynchronous event processing for scalability:

```mermaid
sequenceDiagram
    participant CH as Command Handler
    participant E as Entity
    participant UoW as Unit of Work
    participant MB as Message Bus
    participant EH as Event Handlers

    CH->>E: Command
    E->>E: Creates Events
    CH->>UoW: Save
    UoW->>MB: Publish Events
    MB->>EH: Route Events
```

#### 4. Repository Pattern

- Interface-based design for testability
- Database abstraction
- Easy to mock for testing

#### 5. Unit of Work Pattern

- Transaction management
- Event collection and publishing
- Repository caching

#### 6. Message Bus Pattern

- Centralized command/event handling
- Type-safe routing
- Async processing

### Module Structure

Each module follows a consistent structure:

```
module/
├── entrypoint/          # HTTP handlers and routes
│   └── handler/         # Request handlers
├── service_layer/       # Application services
│   ├── command_handler/ # Write operations
│   └── event_handler/   # Event processing
├── domain/              # Business logic
│   ├── entity/         # Domain entities
│   ├── commands/        # Command DTOs
│   └── events/          # Domain events
├── query/               # Read operations (CQRS)
├── adapter/             # Infrastructure adapters
│   ├── repository/      # Data access
│   └── migrations/      # Database migrations
└── bootstrap.go         # Module initialization
```

### Main Modules

#### 👤 Account Module

- User registration and authentication
- JWT token management
- User profiles with avatar generation
- Session management

#### 🛍️ Products Module

- Product management (CRUD)
- Category management
- Product reviews and ratings
- Product aggregates (features, details, specs)
- Image attachments

---

## 📁 Project Structure

```
backend/
├── 📂 cmd/                    # Application entry points
│   ├── commands/              # CLI commands
│   │   ├── http.go           # HTTP server command
│   │   ├── migrate.go        # Migration commands
│   │   └── root.go           # Root command
│   └── main.go               # Main entry point
│
├── 📂 config/                 # Configuration files
│   ├── config-development.yml # Development config
│   ├── config-docker.yml     # Docker config
│   ├── config-production.yml  # Production config
│   └── config.go             # Config loader
│
├── 📂 internal/               # Application code
│   ├── account/              # User management module
│   │   ├── adapter/          # Infrastructure adapters
│   │   ├── domain/           # Domain layer
│   │   ├── entrypoint/       # HTTP handlers
│   │   ├── query/            # Read operations
│   │   └── service_layer/    # Application services
│   └── products/             # Product management module
│       ├── adapter/          # Infrastructure adapters
│       ├── domain/           # Domain layer
│       ├── entrypoint/       # HTTP handlers
│       ├── query/            # Read operations
│       └── service_layer/   # Application services
│
├── 📂 pkg/                    # Reusable packages
│   └── framework/            # Framework components
│       ├── adapter/          # Base adapters
│       ├── api/              # API utilities
│       ├── errors/           # Error handling
│       ├── helpers/          # Helper functions
│       ├── infrastructure/   # Infrastructure services
│       └── service_layer/    # Service layer utilities
│
├── 📂 docker/                 # Docker configurations
│   ├── docker-compose.yml    # Multi-container setup
│   ├── prometheus/           # Prometheus config
│   ├── grafana/             # Grafana config
│   ├── elk/                 # ELK stack config
│   └── redis/               # Redis config
│
├── 📂 docs/                   # API documentation
│   ├── swagger.json          # Swagger JSON
│   ├── swagger.yaml          # Swagger YAML
│   └── apwp_aa01.png         # Architecture diagram
│
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── Makefile                   # Build automation
└── Dockerfile                # Container definition
```

---

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.25 or higher
- **PostgreSQL** 12 or higher
- **Redis** 6 or higher
- **Docker & Docker Compose** (optional, for full stack)

### Installation

#### 1️⃣ Clone the Repository

```bash
git clone <repository-url>
cd shikposh/backend
```

#### 2️⃣ Install Dependencies

```bash
go mod download
```

#### 3️⃣ Configure the Application

Edit configuration files in the `config/` directory:

```yaml
# config/config-development.yml
database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: shikposh

redis:
  host: localhost
  port: 6379
```

#### 4️⃣ Run Database Migrations

```bash
# Using Make
make migrate-up

# Or directly
go run cmd/main.go migrate up
```

#### 5️⃣ Start the Server

```bash
# Using Make
make run

# Or directly
go run cmd/main.go http
```

The server will start on `http://localhost:8000` (default port).

### 🐳 Docker Setup

For a complete development environment with all services:

```bash
cd docker
docker-compose up -d
```

This starts:

- PostgreSQL database
- Redis cache
- Prometheus metrics
- Grafana dashboards
- ELK stack for logging
- **Jaeger** for distributed tracing
- Kafka for event streaming

---

## 📜 Available Commands

### Development

```bash
# Run the HTTP server
make run
go run cmd/main.go http

# Run with custom config
go run cmd/main.go http --config config/config-development.yml
```

### Database Migrations

```bash
# Run migrations
make migrate-up
go run cmd/main.go migrate up

# Rollback migrations
make migrate-down
go run cmd/main.go migrate down
```

### Testing

```bash
# Run all tests
make test
go test ./tests/... -v

# Run integration tests
make test-integration
TEST_TYPE=integration go test ./tests/integration/... -v
```

### API Documentation

```bash
# Generate Swagger documentation
make swagger
swag fmt && swag init -g ./cmd/main.go -o ./docs
```

---

## 📚 API Documentation

### Swagger UI

Interactive API documentation is available at:

- **Swagger UI**: `http://localhost:8000/swagger/index.html`
- **Swagger JSON**: `http://localhost:8000/swagger.json`
- **Swagger YAML**: `http://localhost:8000/swagger.yaml`

### Main API Endpoints

#### 🔐 Authentication

| Method | Endpoint                  | Description       |
| ------ | ------------------------- | ----------------- |
| `POST` | `/api/v1/public/register` | User registration |
| `POST` | `/api/v1/public/login`    | User login        |
| `POST` | `/api/v1/public/logout`   | User logout       |

#### 🛍️ Products

| Method | Endpoint                                     | Description                      |
| ------ | -------------------------------------------- | -------------------------------- |
| `GET`  | `/api/v1/public/products`                    | List all products (with filters) |
| `GET`  | `/api/v1/public/products/:slug`              | Get product by slug              |
| `GET`  | `/api/v1/public/products/featured`           | Get featured products            |
| `GET`  | `/api/v1/public/products/category/:category` | Get products by category         |

**Query Parameters:**

- `q` - Search query
- `category` - Category slug
- `min` - Minimum price
- `max` - Maximum price
- `rating` - Minimum rating
- `featured` - Featured products only
- `tags` - Comma-separated tags
- `sort` - Sort order (price_asc, price_desc, rating, newest)

#### 📂 Categories

| Method | Endpoint                    | Description         |
| ------ | --------------------------- | ------------------- |
| `GET`  | `/api/v1/public/categories` | List all categories |

#### ⭐ Reviews

| Method  | Endpoint                              | Description                 |
| ------- | ------------------------------------- | --------------------------- |
| `GET`   | `/api/v1/public/products/:id/reviews` | Get product reviews         |
| `POST`  | `/api/v1/public/reviews`              | Create a review             |
| `PATCH` | `/api/v1/public/reviews/:id`          | Update review helpful count |

#### 👤 User Profile

| Method | Endpoint                          | Description      |
| ------ | --------------------------------- | ---------------- |
| `GET`  | `/api/v1/public/users/:id`        | Get user profile |
| `GET`  | `/api/v1/public/users/:id/avatar` | Get user avatar  |

---

## 🔒 Security Features

### Authentication & Authorization

- **JWT Tokens** - Secure token-based authentication
- **bcrypt Hashing** - Industry-standard password hashing (cost: 10)
- **Token Expiration** - Configurable token expiration
- **Session Management** - Redis-based session storage

### Input Validation

- Request validation using Fiber validators
- SQL injection prevention via GORM
- XSS protection in error messages
- Secure error handling (no sensitive data leakage)

### Best Practices

- Environment-based configuration
- Secure default settings
- HTTPS support in production
- CORS configuration
- Rate limiting (configurable)

---

## 📊 Monitoring & Observability

### Metrics (Prometheus)

The application exposes metrics at `/metrics`:

- HTTP request duration
- Request count by endpoint
- Error rates
- Database query performance
- Cache hit/miss rates

### Logging (ELK Stack)

Structured logging with Zerolog:

- **Elasticsearch** - Log storage and indexing
- **Filebeat** - Log collection agent
- **Kibana** - Log visualization and analysis

### Dashboards (Grafana)

Pre-configured dashboards for:

- Application performance
- Database metrics
- Cache performance
- Error tracking
- Request patterns

### Distributed Tracing (Jaeger)

**Jaeger** integration via OpenTelemetry for:

- **Request Tracing** - End-to-end request tracing across services
- **Performance Analysis** - Identify bottlenecks and slow operations
- **Service Dependencies** - Visualize service interactions
- **Span Analysis** - Detailed span timing and metadata
- **Trace Search** - Search and filter traces by tags and attributes

**Access Jaeger UI:**

- **Jaeger UI**: `http://localhost:16686`
- **OTLP HTTP Endpoint**: `http://localhost:4318`
- **OTLP gRPC Endpoint**: `http://localhost:4317`

**Features:**

- OpenTelemetry (OTLP) protocol support
- Configurable sampling rates
- Service and environment tagging
- Trace context propagation

---

## ⚡ Performance Optimizations

### Database

- **Connection Pooling** - Optimized connection management
- **Query Optimization** - Efficient GORM queries
- **Indexes** - Strategic database indexes
- **Prepared Statements** - SQL injection prevention + performance

### Caching

- **Redis Caching** - Frequently accessed data
- **Cache Invalidation** - Smart cache invalidation strategies
- **TTL Management** - Configurable cache expiration

### Concurrency

- **Goroutines** - Concurrent request processing
- **Channel-based Communication** - Efficient event handling
- **Async Event Processing** - Non-blocking operations

### Code Optimizations

- **Zero-copy** where possible
- **Efficient serialization** (JSON)
- **Minimal allocations** in hot paths
- **Connection reuse** for external services

---

## 🧪 Testing

### Running Tests

```bash
# Run all unit tests
make test

# Run integration tests
make test-integration

# Run with coverage
go test ./... -cover
```

### Test Structure

```
tests/
├── unit/              # Unit tests
└── integration/       # Integration tests
```

### Testing Best Practices

- Unit tests for business logic
- Integration tests for API endpoints
- Mock repositories for isolation
- Test fixtures for consistent data

---

## 🚀 Deployment

### Production Build

```bash
# Build binary
go build -o bin/shikposh cmd/main.go

# Run production server
./bin/shikposh http --config config/config-production.yml
```

### Docker Deployment

```bash
# Build Docker image
docker build -t shikposh-backend .

# Run container
docker run -p 8000:8000 \
  -e DATABASE_URL=postgres://... \
  -e REDIS_URL=redis://... \
  shikposh-backend
```

### Environment Variables

```env
# Database
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=password
DATABASE_NAME=shikposh

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# Server
SERVER_PORT=8000
SERVER_HOST=0.0.0.0

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Jaeger Tracing
JAEGER_ENABLED=true
JAEGER_SERVICE_NAME=shikposh-backend
JAEGER_ENVIRONMENT=development
JAEGER_OTLP_ENDPOINT=http://localhost:4318
JAEGER_SAMPLING_RATE=1.0
```

---

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. 🍴 Fork the repository
2. 🌿 Create a feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🔀 Open a Pull Request

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Write comprehensive tests
- Update documentation

---

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## 👨‍💻 Developer

**Ali Mahdavi**

- GitHub: [@ali-mahdavi-dev](https://github.com/ali-mahdavi-dev)

---

<div align="center">

**Built with ❤️ to showcase enterprise backend development skills**

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v3-00ADD8)](https://gofiber.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12+-336791?logo=postgresql)](https://www.postgresql.org/)

⭐ If you find this project interesting, give it a Star!

</div>
