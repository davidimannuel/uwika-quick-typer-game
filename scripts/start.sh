#!/bin/bash

echo "🚀 Starting Quick Typer Development Environment..."

# Check if .env file exists
if [ ! -f .env ]; then
    echo "⚠️  .env file not found, copying from .env.example..."
    cp .env.example .env
fi

# Start docker-compose
docker-compose up --build

echo "✅ Environment started!"
echo "📝 Backend API: http://localhost:8080"
echo "🎨 Admin Panel: http://localhost:3000"
echo "🗄️  PostgreSQL: localhost:5432"

