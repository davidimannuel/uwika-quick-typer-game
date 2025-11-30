#!/bin/bash

echo "🚀 Starting Quick Typer services locally..."

# Check if PostgreSQL is running
if ! docker ps | grep -q quick_typer_db; then
    echo "Starting PostgreSQL..."
    docker run -d \
      --name quick_typer_db \
      -e POSTGRES_USER=postgres \
      -e POSTGRES_PASSWORD=s3cret \
      -e POSTGRES_DB=quick_typer \
      -p 5432:5432 \
      postgres:16-alpine
    
    echo "Waiting for PostgreSQL to be ready..."
    sleep 5
fi

# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=s3cret
export DB_NAME=quick_typer
export DB_SSLMODE=disable

# Start API in background
echo "Starting API on port 8080..."
PORT=8080 go run cmd/api/main.go &
API_PID=$!

# Start Admin Web in background
echo "Starting Admin Web on port 3000..."
PORT=3000 go run cmd/admin-web/main.go &
ADMIN_PID=$!

echo ""
echo "✅ Services started!"
echo "📝 API: http://localhost:8080"
echo "🎨 Admin Panel: http://localhost:3000"
echo ""
echo "Press Ctrl+C to stop all services..."

# Wait for interrupt
trap "kill $API_PID $ADMIN_PID; exit" INT
wait

