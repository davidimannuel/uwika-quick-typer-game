#!/bin/bash

echo "🏗️  Building backend services..."

# Create bin directory if not exists
mkdir -p bin

# Build API
echo "Building API..."
CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api
echo "✅ API built successfully!"

# Build Admin Web
echo "Building Admin Web..."
CGO_ENABLED=0 GOOS=linux go build -o bin/admin-web ./cmd/admin-web
echo "✅ Admin Web built successfully!"

echo ""
echo "✅ All builds completed!"
echo "📦 Binaries:"
echo "   - bin/api"
echo "   - bin/admin-web"

