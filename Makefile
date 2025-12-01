# https://github.com/golang-migrate/migrate
# make create-db-migration name=<migration_name>, example: make create-db-migration name=create_users_table
create-db-migration:
	migrate create -ext sql -dir dbscript/migrations $(name)

# Local migrations (requires migrate CLI installed)
migrate-up:
	migrate -source file://dbscript/migrations -database postgres://postgres:s3cret@127.0.0.1:5432/quick_typer?sslmode=disable up

migrate-down:
	migrate -source file://dbscript/migrations -database postgres://postgres:s3cret@127.0.0.1:5432/quick_typer?sslmode=disable down -all

migrate-force:
	migrate -source file://dbscript/migrations -database postgres://postgres:s3cret@127.0.0.1:5432/quick_typer?sslmode=disable force $(version)

# Docker migrations (no CLI required)
docker-migrate-up:
	docker run --rm -v $(PWD)/dbscript/migrations:/migrations --network host migrate/migrate:v4.17.0 \
		-path=/migrations/ -database postgres://postgres:s3cret@localhost:5432/quick_typer?sslmode=disable up

docker-migrate-down:
	docker run --rm -v $(PWD)/dbscript/migrations:/migrations --network host migrate/migrate:v4.17.0 \
		-path=/migrations/ -database postgres://postgres:s3cret@localhost:5432/quick_typer?sslmode=disable down -all

docker-migrate-force:
	docker run --rm -v $(PWD)/dbscript/migrations:/migrations --network host migrate/migrate:v4.17.0 \
		-path=/migrations/ -database postgres://postgres:s3cret@localhost:5432/quick_typer?sslmode=disable force $(version)

# Build commands
build-api:
	CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api

build-admin:
	CGO_ENABLED=0 GOOS=linux go build -o bin/admin-web ./cmd/admin-web

build-all: build-api build-admin

# Run local development
run-api:
	go run cmd/api/main.go

run-admin:
	go run cmd/admin-web/main.go

run-dev:
	./scripts/dev.sh

# Docker commands
run-all:
	docker-compose up -d

stop-all:
	docker-compose down

restart-all:
	docker-compose down
	docker-compose up --build -d

run-all-build:
	docker-compose up --build -d

# Run database only (for local API development)
run-db:
	docker-compose up -d postgres migrate

stop-db:
	docker-compose stop postgres migrate

# Local development workflow
dev-local:
	@echo "Starting database..."
	@make run-db
	@echo "Waiting for database to be ready..."
	@sleep 5
	@echo "Database ready! Now run 'make run-api' in another terminal"
	@echo "To expose API with ngrok, run 'make ngrok' in another terminal"

logs:
	docker-compose logs -f

logs-api:
	docker-compose logs -f api

logs-migrate:
	docker-compose logs migrate

logs-db:
	docker-compose logs -f postgres

# Database commands
db-reset:
	docker-compose down -v
	docker-compose up --build -d

# Utility
# 8080 is the port of the API
ngrok:
	ngrok http 8080

# Complete local development setup
help-local:
	@echo "=== Local Development Guide ==="
	@echo ""
	@echo "1. Start database (Postgres + migrations):"
	@echo "   make run-db"
	@echo ""
	@echo "2. Run API locally (in new terminal):"
	@echo "   make run-api"
	@echo ""
	@echo "3. Expose API with ngrok (in new terminal):"
	@echo "   make ngrok"
	@echo ""
	@echo "4. Stop database:"
	@echo "   make stop-db"
	@echo ""
	@echo "Database connection details:"
	@echo "  Host: localhost"
	@echo "  Port: 5432"
	@echo "  User: postgres"
	@echo "  Password: s3cret"
	@echo "  Database: quick_typer"