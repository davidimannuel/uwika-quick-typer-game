# Quick Typer Backend & Admin Web

Backend Golang dengan Clean Architecture untuk game Quick Typer, lengkap dengan Admin Web Panel.

## 🏗️ Arsitektur

Proyek ini menggunakan **Clean Architecture** dengan struktur:

```
uwika_quick_typer_game/
├── cmd/
│   ├── api/                    # REST API Server
│   └── admin-web/              # Admin Web Server
├── internal/
│   ├── domain/                 # Business Logic Layer
│   │   ├── models/            # Domain entities (tanpa JSON tags)
│   │   └── repositories/      # Repository interfaces (Ports)
│   ├── application/           # Application Layer
│   │   └── services/          # Business logic & use cases
│   └── infrastructure/        # Infrastructure Layer
│       ├── database/          # Database connection
│       ├── http/              # HTTP layer
│       │   ├── dto/           # Request/Response models (dengan JSON tags)
│       │   ├── handlers/      # HTTP handlers
│       │   ├── middleware/    # Middleware
│       │   └── router/        # Router setup
│       └── persistence/       # Database implementations (Adapters)
│           └── postgres/      # PostgreSQL repositories
├── dbscript/
│   ├── migrations/            # Database migrations
│   └── create_admin.sh       # Script create admin user
├── scripts/                   # Helper scripts
├── go.mod
├── docker-compose.yml
├── Dockerfile.api
└── Dockerfile.admin
```

### Principles

- **Domain-Driven Design (DDD)**: Domain models bersih tanpa dependency external
- **Port & Adapter**: Repository interfaces di domain, implementasi di infrastructure
- **Dependency Inversion**: High-level modules tidak depend pada low-level modules
- **Clean Architecture**: Separation of concerns yang jelas
- **Microservices**: API dan Admin Web terpisah sebagai dua services berbeda

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.23+ (untuk development lokal)
- Make (optional)

### 1. Clone & Setup

```bash
# Clone repository
cd uwika_quick_typer_game

# Copy environment variables (optional)
cp .env.example .env
```

### 2. Start dengan Docker

```bash
# Start semua services (PostgreSQL, API, Admin Web)
make run-all

# Atau gunakan docker-compose langsung
docker-compose up --build -d
```

Services yang berjalan:
- **REST API**: http://localhost:3000
- **Admin Web**: http://localhost:3000
- **PostgreSQL**: localhost:5432

### 3. Migrasi Database

Migrasi otomatis running saat PostgreSQL container start pertama kali. Atau jalankan manual:

```bash
make migrate-up
```

### 4. Buat Admin User

```bash
cd dbscript
chmod +x create_admin.sh
./create_admin.sh
```

Default admin credentials:
- Username: `admin`
- Password: `admin123`

## 📦 Development Lokal (Tanpa Docker)

### 1. Install Dependencies

```bash
go mod download
```

### 2. Setup Database

```bash
# Start PostgreSQL dengan Docker
docker run -d \
  --name quick_typer_db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=s3cret \
  -e POSTGRES_DB=quick_typer \
  -p 5432:5432 \
  postgres:16-alpine

# Run migrations
make migrate-up
```

### 3. Run Services

**Option A: Run semua dengan script**
```bash
make run-dev
# atau
./scripts/dev.sh
```

**Option B: Run manual (terminal terpisah)**

Terminal 1 - API:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=s3cret
export DB_NAME=quick_typer
export DB_SSLMODE=disable

make run-api
# atau
go run cmd/api/main.go
```

Terminal 2 - Admin Web:
```bash
export PORT=3000

make run-admin
# atau
go run cmd/admin-web/main.go
```

## 🔌 API Endpoints

### Authentication

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/auth/register` | POST | Register user baru |
| `/api/auth/login` | POST | Login dan dapatkan token |
| `/api/auth/profile` | GET | Get user profile (perlu auth) |

### Game API (Require User Token)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/stages` | GET | List semua stages aktif |
| `/api/stage/:id` | GET | Detail stage dengan phrases |
| `/api/score/submit` | POST | Submit score permainan |
| `/api/leaderboard` | GET | Get leaderboard by stage |

### Admin API (Require Admin Token)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/admin/stage` | POST | Buat stage baru |
| `/admin/stage/:id` | PUT | Update stage |
| `/admin/stage/:id` | DELETE | Hapus stage |
| `/admin/stages` | GET | List semua stages |
| `/admin/phrase` | POST | Buat phrase baru |
| `/admin/phrase/:id` | PUT | Update phrase |
| `/admin/phrase/:id` | DELETE | Hapus phrase |
| `/admin/phrases` | GET | List phrases by stage |

## 🎨 Admin Web Panel

Access: http://localhost:3000

Features:
- ✅ Login dengan admin credentials
- ✅ Manage Stages (Create, Update, Delete)
- ✅ Manage Phrases (Create, Delete)
- ✅ Modern & Responsive UI
- ✅ Real-time data updates
- ✅ Embedded static files (tidak perlu nginx)

## 🛠️ Makefile Commands

```bash
# Database migrations
make create-db-migration name=migration_name
make migrate-up
make migrate-down

# Build commands
make build-api              # Build API binary
make build-admin            # Build Admin Web binary
make build-all              # Build semua

# Run local development
make run-api                # Run API saja
make run-admin              # Run Admin Web saja
make run-dev                # Run semua (API + Admin + DB)

# Docker commands
make run-all                # Start all services
make stop-all               # Stop all services  
make restart-all            # Restart
make run-all-build          # Rebuild and start

# Utility
make ngrok                  # Expose local server via ngrok
```

## 📚 Dependencies

### Backend
- **gin-gonic/gin**: HTTP web framework
- **lib/pq**: PostgreSQL driver
- **google/uuid**: UUID generation
- **golang.org/x/crypto**: Bcrypt hashing

### Embedded
- Admin web menggunakan Go embed untuk static files
- Tidak perlu nginx atau web server terpisah

## 🐛 Troubleshooting

### Database connection error

```bash
# Check PostgreSQL is running
docker ps | grep postgres

# Check connection
psql -h localhost -p 5432 -U postgres -d quick_typer
```

### Port already in use

```bash
# Check ports
lsof -i :8080  # API
lsof -i :3000  # Admin Web
lsof -i :5432  # PostgreSQL

# Kill process
kill -9 <PID>
```

### Migration errors

```bash
# Reset database
make migrate-down
make migrate-up
```

## 📄 Project Structure

```
uwika_quick_typer_game/
├── cmd/
│   ├── api/                    # REST API (Port 8080)
│   │   └── main.go
│   └── admin-web/              # Admin Web (Port 3000)
│       ├── main.go
│       └── static/             # Embedded static files
│           ├── index.html
│           └── app.js
├── internal/
│   ├── domain/                 # Domain Layer
│   ├── application/            # Application Layer
│   └── infrastructure/         # Infrastructure Layer
├── dbscript/
│   └── migrations/
├── scripts/
├── Dockerfile.api              # API container
├── Dockerfile.admin            # Admin Web container
├── docker-compose.yml          # Multi-container setup
└── Makefile
```

## 🔒 Security

- Password hashing menggunakan bcrypt
- Token authentication dengan SHA-256
- Token expiry 30 hari
- Middleware untuk autentikasi & autorisasi
- CORS enabled untuk development
- Role-based access control (User & Admin)

---

**Quick Typer Backend** - Built with ❤️ using Go & Clean Architecture

