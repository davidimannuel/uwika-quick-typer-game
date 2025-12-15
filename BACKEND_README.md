# Quick Typer Backend & Admin Web

Backend Golang dengan Clean Architecture untuk game Quick Typer, lengkap dengan Admin Web Panel.

## 🏗️ Arsitektur

Proyek ini menggunakan **Clean Architecture** dengan struktur:

```
backend/
├── cmd/
│   └── api/                    # Entry point aplikasi
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
```

### Principles

- **Domain-Driven Design (DDD)**: Domain models bersih tanpa dependency external
- **Port & Adapter**: Repository interfaces di domain, implementasi di infrastructure
- **Dependency Inversion**: High-level modules tidak depend pada low-level modules
- **Clean Architecture**: Separation of concerns yang jelas

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.23+ (untuk development lokal)
- Make (optional)

### 1. Clone & Setup

```bash
# Clone repository
cd uwika_quick_typer_game

# Copy environment variables
cp .env.example .env
```

### 2. Start dengan Docker

```bash
# Start semua services (PostgreSQL, Backend, Admin Web)
make run-all

# Atau gunakan docker-compose langsung
docker-compose up --build -d
```

Services yang berjalan:
- **Backend API**: http://localhost:3000
- **Admin Web**: http://localhost:3000
- **PostgreSQL**: localhost:5432

### 3. Migrasi Database

```bash
# Jalankan migrasi
make migrate-up

# Atau gunakan migrate tool langsung
migrate -source file://dbscript/migrations \
  -database postgres://postgres:s3cret@localhost:5432/quick_typer?sslmode=disable up
```

### 4. Buat Admin User

```bash
# Jalankan script create admin
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
cd backend
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
```

### 3. Jalankan Migrasi

```bash
make migrate-up
```

### 4. Run Backend

```bash
# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=s3cret
export DB_NAME=quick_typer
export DB_SSLMODE=disable
export PORT=8080

# Run
cd backend
go run cmd/api/main.go
```

### 5. Run Admin Web

```bash
# Serve dengan simple HTTP server
cd admin-web
python3 -m http.server 3000

# Atau gunakan npx
npx serve -p 3000
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

## 🧪 Testing API

### 1. Register User

```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### 2. Login

```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

Response:
```json
{
  "user_id": "xxx",
  "access_token": "yyy",
  "token_expires_at": "2024-01-01T00:00:00Z"
}
```

### 3. Get Stages (Dengan Token)

```bash
curl http://localhost:3000/api/stages \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 4. Submit Score

```bash
curl -X POST http://localhost:3000/api/score/submit \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "stage_id": "stage-001",
    "total_time_ms": 15000,
    "total_errors": 2
  }'
```

## 🎨 Admin Web Panel

Access: http://localhost:3000

Features:
- ✅ Login dengan admin credentials
- ✅ Manage Stages (Create, Update, Delete)
- ✅ Manage Phrases (Create, Delete)
- ✅ Modern & Responsive UI
- ✅ Real-time data updates

## 📝 Database Schema

### Users
- `user_id` (PK)
- `username` (Unique)
- `password_hash`
- `role` (user/admin)

### PersonalAccessTokens
- `id` (PK)
- `user_id` (FK → users)
- `token` (Indexed)
- `expires_at`
- `revoked_at`

### Stages
- `stage_id` (PK)
- `name`
- `theme`
- `difficulty` (easy/medium/hard)
- `is_active`

### Phrases
- `phrase_id` (PK)
- `stage_id` (FK → stages)
- `text`
- `sequence_number`
- `base_multiplier`

### Scores
- `user_id` (FK → users) (Composite PK)
- `stage_id` (FK → stages) (Composite PK)
- `final_score`
- `total_time_ms`
- `total_errors`

## 🧮 Score Calculation

Formula sesuai README:

```
Final Score = (Σ(phrase_length × multiplier) / time_in_seconds) - (errors × 50)
```

Implementasi di `backend/internal/application/services/game_service.go`

## 🛠️ Makefile Commands

```bash
# Database migrations
make create-db-migration name=migration_name
make migrate-up
make migrate-down

# Docker commands
make run-all              # Start all services
make stop-all             # Stop all services  
make restart-all          # Restart all services
make run-all-build        # Rebuild and start

# Development
make ngrok                # Expose local server via ngrok
```

## 📂 Project Structure

```
uwika_quick_typer_game/
├── backend/               # Backend source code
│   ├── cmd/
│   │   └── api/          # Main application
│   └── internal/
│       ├── domain/       # Domain layer
│       ├── application/  # Application layer
│       └── infrastructure/ # Infrastructure layer
├── admin-web/            # Admin web interface
│   ├── index.html
│   └── app.js
├── dbscript/
│   ├── migrations/       # Database migrations
│   └── create_admin.sh   # Admin user script
├── scripts/              # Helper scripts
├── Dockerfile            # Backend container
├── docker-compose.yml    # Multi-container setup
├── nginx.conf            # Nginx config for admin web
└── Makefile
```

## 🔒 Security

- Password hashing menggunakan bcrypt
- Token authentication dengan SHA-256
- Token expiry 30 hari
- Middleware untuk autentikasi & autorisasi
- CORS enabled untuk development

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
lsof -i :8080  # Backend
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

## 📚 Dependencies

### Backend
- **gin-gonic/gin**: HTTP web framework
- **lib/pq**: PostgreSQL driver
- **google/uuid**: UUID generation
- **golang.org/x/crypto**: Bcrypt hashing

### Tools
- **golang-migrate/migrate**: Database migrations
- **nginx**: Web server untuk admin panel

## 👨‍💻 Development

Struktur project ini mengikuti Clean Architecture principles:

1. **Domain Layer**: Pure business logic, tidak ada dependencies
2. **Application Layer**: Use cases dan orchestration
3. **Infrastructure Layer**: External dependencies (DB, HTTP)

Tidak menggunakan factory patterns sesuai preferensi [[memory:6380932]].

## 📄 License

MIT License - Feel free to use for educational purposes.

---

**Quick Typer Backend** - Built with ❤️ using Go & Clean Architecture

