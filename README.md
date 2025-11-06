# 🚀 Shikposh Backend

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Architecture](https://img.shields.io/badge/Architecture-Clean%20Architecture-blueviolet.svg)
![Pattern](https://img.shields.io/badge/Pattern-DDD%20%7C%20CQRS%20%7C%20Event%20Sourcing-orange.svg)

**یک Backend API مدرن و مقیاس‌پذیر با معماری Clean Architecture و Domain-Driven Design**

[Features](#-features) • [Architecture](#-architecture) • [Tech Stack](#-tech-stack) • [Getting Started](#-getting-started) • [Project Structure](#-project-structure)

</div>

---

## 📋 فهرست مطالب

- [معرفی](#-معرفی)
- [ویژگی‌ها](#-ویژگی‌ها)
- [معماری](#-معماری)
- [تکنولوژی‌ها](#-تکنولوژی‌ها)
- [ساختار پروژه](#-ساختار-پروژه)
- [شروع کار](#-شروع-کار)
- [API Documentation](#-api-documentation)
- [Monitoring & Observability](#-monitoring--observability)

---

## 🎯 معرفی

**Shikposh Backend** یک پروژه Backend API مدرن است که با استفاده از **Clean Architecture** و **Domain-Driven Design (DDD)** پیاده‌سازی شده است. این پروژه از الگوهای پیشرفته‌ای مانند **CQRS**, **Event Sourcing**, و **Message Bus** برای ایجاد یک سیستم مقیاس‌پذیر و قابل نگهداری استفاده می‌کند.

### ✨ ویژگی‌های کلیدی

- 🏗️ **Clean Architecture** - جداسازی کامل لایه‌ها و وابستگی‌ها
- 🎯 **Domain-Driven Design** - طراحی بر اساس Domain Model
- 📨 **CQRS Pattern** - جداسازی Command و Query
- 🎪 **Event-Driven Architecture** - استفاده از Domain Events
- 🔄 **Unit of Work Pattern** - مدیریت Transaction و Event Collection
- 🚌 **Message Bus** - پردازش Asynchronous Commands و Events
- 🔒 **Thread-Safe Operations** - جلوگیری از Race Conditions
- 🛡️ **Graceful Shutdown** - خاموش شدن امن و کنترل شده

---

## 🏗️ معماری

### Clean Architecture Layers

```
┌─────────────────────────────────────────────────┐
│           Presentation Layer (HTTP)             │
│         (Fiber Router, Handlers)                │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│          Application Layer                      │
│    (Command Handlers, Event Handlers)           │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│            Domain Layer                         │
│    (Entities, Commands, Events, Business Logic) │
└─────────────────────────────────────────────────┘
                      ↓
┌─────────────────────────────────────────────────┐
│         Infrastructure Layer                    │
│  (Database, Cache, Message Bus, Logging)         │
└─────────────────────────────────────────────────┘
```

### الگوهای طراحی استفاده شده

#### 1. **Domain-Driven Design (DDD)**

- **Entities**: User, Profile, Token
- **Commands**: RegisterUser, LoginUser, Logout
- **Domain Events**: RegisterUserEvent
- **Value Objects**: در صورت نیاز

#### 2. **CQRS (Command Query Responsibility Segregation)**

- **Commands**: برای تغییر state (Register, Login, Logout)
- **Queries**: برای خواندن داده (GetUser, ViewUser)

#### 3. **Event-Driven Architecture**

- Domain Events برای decoupling
- Event Handlers برای side effects
- Nested Events Support

#### 4. **Unit of Work Pattern**

- مدیریت Transaction
- جمع‌آوری Domain Events
- Repository Caching per Transaction

#### 5. **Message Bus Pattern**

- Centralized Command/Event Handling
- Async Event Processing
- Graceful Shutdown Support

---

## 🛠️ تکنولوژی‌ها

### Core Technologies

- **Go 1.25** - زبان برنامه‌نویسی
- **Fiber v3** - Web Framework (بر پایه FastHTTP)
- **GORM** - ORM برای PostgreSQL
- **PostgreSQL** - Database اصلی
- **Redis** - Caching و Session Management

### Infrastructure & Tools

- **Docker & Docker Compose** - Containerization
- **Prometheus** - Metrics Collection
- **Grafana** - Monitoring Dashboards
- **ELK Stack** - Logging (Elasticsearch, Logstash, Kibana)
- **Filebeat** - Log Shipper
- **Alertmanager** - Alerting

### Libraries & Frameworks

- **Zerolog** - Structured Logging
- **JWT** - Authentication
- **Swagger** - API Documentation
- **Cobra** - CLI Framework
- **Viper** - Configuration Management
- **Kafka** - Message Queue (برای Event Streaming)
- **Socket.IO** - WebSocket Support

---

## 📁 ساختار پروژه

```
backend/
├── cmd/                          # Application Entry Points
│   ├── main.go                   # Main entry point
│   └── commands/                 # CLI Commands
│       ├── http.go               # HTTP server command
│       ├── migrate.go            # Migration command
│       └── root.go               # Root command
│
├── internal/                     # Private Application Code
│   └── account/                  # Account Module (Domain)
│       ├── adapter/              # Infrastructure Adapters
│       │   ├── migrations/       # Database Migrations
│       │   ├── repository/       # Repository Implementations
│       │   └── avatar_generator.go
│       ├── domain/               # Domain Layer
│       │   ├── commands/         # Command DTOs
│       │   ├── entity/           # Domain Entities
│       │   └── events/           # Domain Events
│       ├── entryporint/          # Presentation Layer
│       │   └── handler/          # HTTP Handlers
│       ├── query/                # Query Handlers (CQRS)
│       └── service_layer/        # Application Services
│           ├── command_handler/  # Command Handlers
│           └── event_handler/    # Event Handlers
│
├── pkg/                          # Public Packages (Reusable)
│   └── framework/                # Framework Components
│       ├── adapter/              # Base Adapters
│       ├── api/                  # API Utilities
│       ├── errors/               # Error Handling
│       ├── infrastructure/       # Infrastructure Services
│       └── service_layer/        # Service Layer Patterns
│           ├── messagebus/       # Message Bus Implementation
│           └── unit_of_work/     # Unit of Work Pattern
│
├── config/                       # Configuration Files
├── docker/                       # Docker Configurations
├── docs/                         # API Documentation
└── Makefile                      # Build Automation
```

### معماری لایه‌ای

1. **Domain Layer** (`internal/*/domain/`)

   - Entities, Commands, Events
   - Business Logic خالص
   - بدون وابستگی به Infrastructure

2. **Application Layer** (`internal/*/service_layer/`)

   - Command Handlers
   - Event Handlers
   - Orchestration Logic

3. **Infrastructure Layer** (`pkg/framework/infrastructure/`)

   - Database Connections
   - Cache (Redis)
   - Message Queue (Kafka)
   - Logging

4. **Presentation Layer** (`internal/*/entryporint/`)
   - HTTP Handlers
   - Request/Response Mapping
   - Validation

---

## 🚀 شروع کار

### پیش‌نیازها

- Go 1.25 یا بالاتر
- PostgreSQL 12+
- Redis 6+
- Docker & Docker Compose (اختیاری)

### نصب و راه‌اندازی

1. **Clone Repository**

```bash
git clone git@github.com:ali-mahdavi-dev/shikposh-backend.git
cd shikposh-backend
```

2. **نصب Dependencies**

```bash
go mod download
```

3. **اجرای Migrations**

```bash
go run cmd/main.go migrate
```

5. **اجرای Server**

```bash
go run cmd/main.go http
```

### استفاده از Docker

```bash
docker-compose up -d
```

---

## 📚 API Documentation

API Documentation با استفاده از Swagger در دسترس است:

- **Swagger UI**: `http://localhost:8000/swagger/index.html`
- **Swagger JSON**: `http://localhost:8000/swagger.json`

### Endpoints اصلی

#### Authentication

- `POST /api/v1/public/register` - ثبت‌نام کاربر
- `POST /api/v1/public/login` - ورود کاربر
- `POST /api/v1/public/logout` - خروج کاربر

#### Health & Monitoring

- `GET /health` - Health Check
- `GET /ready` - Readiness Check
- `GET /metrics` - Prometheus Metrics

---

## 📊 Monitoring & Observability

### Metrics (Prometheus)

- HTTP Request Metrics
- Database Connection Pool Metrics
- Custom Business Metrics

### Logging (ELK Stack)

- Structured Logging با Zerolog
- Centralized Log Management
- Log Aggregation و Analysis

### Dashboards (Grafana)

- Application Performance
- Database Metrics
- System Resources

---

## 🔐 Security Features

- **JWT Authentication** - Token-based Authentication
- **Password Hashing** - bcrypt برای Hash کردن Passwords
- **Input Validation** - Validation در تمام Endpoints
- **Error Handling** - Error Messages امن و بدون اطلاعات حساس

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...
```

---

## 🏭 Deployment

### Build

```bash
make build
```

### Docker Build

```bash
docker build -t shikposh-backend .
```

### Production Deployment

```bash
# با استفاده از Docker Compose
docker-compose -f docker-compose.prod.yml up -d
```

---

## 📈 Performance

- **Concurrent Request Handling** - پردازش همزمان درخواست‌ها
- **Connection Pooling** - مدیریت اتصالات دیتابیس
- **Caching Strategy** - استفاده از Redis برای Cache
- **Async Event Processing** - پردازش Asynchronous Events

---

## 🤝 Contributing

این پروژه برای نمایش مهارت‌های من در معماری نرم‌افزار و Go Development ساخته شده است.

---

## 📝 License

این پروژه تحت مجوز MIT منتشر شده است.

---

## 💡 مثال‌های معماری

### Domain Event Flow

```go
// 1. User Entity ایجاد می‌شود و Event را اضافه می‌کند
user := entity.NewUser(...)  // RegisterUserEvent اضافه می‌شود

// 2. User ذخیره می‌شود
uow.User(ctx).Save(ctx, user)

// 3. Unit of Work Events را جمع‌آوری می‌کند
uow.CollectNewEvents(ctx, eventCh)

// 4. Message Bus Event را پردازش می‌کند
bus.HandleEvent(ctx, event)

// 5. Event Handler Profile ایجاد می‌کند
eventHandler.RegisterEvent(ctx, event)
```

### Command Handler Pattern

```go
// Command Definition
type RegisterUser struct {
    UserName string `json:"user_name"`
    Email    string `json:"email"`
    // ...
}

// Command Handler
func (h *UserHandler) RegisterHandler(ctx context.Context, cmd *commands.RegisterUser) (*RegisterResult, error) {
    return h.uow.Do(ctx, func(ctx context.Context) error {
        // Business Logic
        user := entity.NewUser(...)
        return h.uow.User(ctx).Save(ctx, user)
    })
}

// Usage via Message Bus
result, err := bus.Handle(ctx, &commands.RegisterUser{...})
```

### Unit of Work Pattern

```go
// Transaction Management
err := uow.Do(ctx, func(ctx context.Context) error {
    user := uow.User(ctx).FindByID(ctx, id)
    profile := uow.Profile(ctx).FindByUserID(ctx, user.ID)
    // All operations in single transaction
    return nil
})

// Event Collection
uow.CollectNewEvents(ctx, eventCh)
```

---

## 👨‍💻 Author

**Ali Mahdavi**

- GitHub: [@ali-mahdavi-dev](https://github.com/ali-mahdavi-dev)

---

<div align="center">

**ساخته شده با ❤️ و Clean Architecture**

⭐ اگر این پروژه برای شما مفید بود، یک Star بدهید!

</div>
