# Quick Typer - Docker Deployment Guide

## 🐳 Docker Setup dengan Migrations

Project ini menggunakan [golang-migrate/migrate](https://github.com/golang-migrate/migrate) untuk database migrations di Docker environment.

### Architecture

```
┌─────────────┐
│  postgres   │ (Port 5432)
└──────┬──────┘
       │ wait healthy
       ↓
┌─────────────┐
│   migrate   │ (run migrations then exit)
└──────┬──────┘
       │ wait completed
       ↓
┌─────────────┐
│     api     │ (Port 8080)
└──────┬──────┘
       │
       ↓
┌─────────────┐
│  admin-web  │ (Port 3000)
└─────────────┘
```

### Services

1. **postgres** - PostgreSQL database
2. **migrate** - Run database migrations (exit after completion)
3. **api** - REST API backend
4. **admin-web** - Admin web interface

## 🚀 Quick Start

### 1. Start All Services
```bash
make run-all-build
# atau
docker-compose up --build -d
```

**Flow:**
1. PostgreSQL starts dan wait sampai healthy
2. Migrate service run migrations dan exit
3. API service start setelah migrations selesai
4. Admin web start setelah API ready

### 2. Check Migration Status
```bash
make logs-migrate
# atau
docker-compose logs migrate
```

Expected output:
```
migrate_1  | 1/u initial_schema (XX.XXs)
migrate_1  | 2/u seed_data (XX.XXs)
```

### 3. Verify Services
```bash
# Check all services
docker-compose ps

# API health check
curl http://localhost:8080/health

# Admin web
open http://localhost:3000
```

## 📋 Migration Commands

### Docker Migrations (No CLI Install Required)

```bash
# Run migrations up
make docker-migrate-up

# Rollback all migrations
make docker-migrate-down

# Force migration to specific version (if stuck)
make docker-migrate-force version=2
```

### Local Migrations (Requires migrate CLI)

Install migrate CLI:
```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Windows
choco install migrate
```

Run migrations:
```bash
# Run migrations up
make migrate-up

# Rollback all migrations
make migrate-down

# Force specific version
make migrate-force version=2
```

## 🔄 Common Workflows

### Fresh Start (Reset Everything)
```bash
# Stop all services dan hapus volumes
docker-compose down -v

# Start fresh
make run-all-build
```

### Update Migrations
```bash
# Create new migration
make create-db-migration name=add_new_table

# Edit migration files in dbscript/migrations/

# Rebuild and apply
make restart-all
```

### Check Migration Logs
```bash
# See migration output
make logs-migrate

# See API logs
make logs-api

# See all logs
make logs
```

### Migration Failed?

If migration stuck or failed:

```bash
# 1. Check current migration version
docker run --rm -v $(PWD)/dbscript/migrations:/migrations --network quick_typer_network migrate/migrate \
  -path=/migrations/ -database postgres://postgres:s3cret@postgres:5432/quick_typer?sslmode=disable version

# 2. Force to specific version
make docker-migrate-force version=1

# 3. Try migrate up again
make docker-migrate-up

# 4. Or reset database completely
make db-reset
```

## 🛠️ Development vs Production

### Development (Local)
```bash
# Run with local migrations
make run-dev
make migrate-up
```

### Production (Docker)
```bash
# Migrations automatic run in docker-compose
make run-all-build
```

## 📁 Migration Files Structure

```
dbscript/
├── migrations/
│   ├── 000001_initial_schema.up.sql    # Create tables
│   ├── 000001_initial_schema.down.sql  # Drop tables
│   ├── 000002_seed_data.up.sql         # Insert seed data
│   └── 000002_seed_data.down.sql       # Delete seed data
└── create_admin.sh                     # Manual admin creation
```

## 🔍 Troubleshooting

### Migration Container Keeps Restarting

Check logs:
```bash
make logs-migrate
```

Common issues:
- Database not ready → Wait for postgres healthy check
- Wrong credentials → Check environment variables
- Migration syntax error → Check SQL syntax in migration files

### API Can't Connect to Database

```bash
# Check if migrate completed successfully
docker-compose ps

# migrate container should be "Exit 0"
# If still running or error, check logs
make logs-migrate
```

### Reset Everything

```bash
# Nuclear option - reset all
make db-reset
```

## 📊 Useful Commands

```bash
# Check services status
docker-compose ps

# View logs
make logs                # All services
make logs-api           # API only
make logs-migrate       # Migration only

# Database operations
make db-reset           # Reset database and restart
make docker-migrate-up  # Manual migration up
make docker-migrate-down # Manual migration down

# Rebuild
make restart-all        # Rebuild and restart all
make run-all-build      # Build and run
```

## 🔐 Environment Variables

Migration menggunakan environment variables dari docker-compose.yml:

```yaml
POSTGRES_USER: postgres
POSTGRES_PASSWORD: s3cret
POSTGRES_DB: quick_typer
```

Untuk production, override dengan `.env` file atau environment variables.

## 📚 References

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate) - Migration tool
- [Migration Best Practices](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)
- Docker Compose Documentation

---

**Note**: Migration service adalah "one-shot" service yang exit setelah selesai. Ini normal behavior dan bukan error.

