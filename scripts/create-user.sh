#!/bin/bash

# Create new user using Docker
# Usage: ./scripts/create-user.sh [username] [password] [role]
# Role: "user" or "admin"

set -e

# Database configuration
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-s3cret}"
DB_NAME="${DB_NAME:-quick_typer}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "👤 Quick Typer - Create User"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# Get username
if [ -n "$1" ]; then
    USERNAME="$1"
else
    read -p "Enter username: " USERNAME
fi

if [ -z "$USERNAME" ]; then
    echo -e "${RED}❌ Error: Username cannot be empty${NC}"
    exit 1
fi

# Get password
if [ -n "$2" ]; then
    NEW_PASSWORD="$2"
else
    read -sp "Enter password: " NEW_PASSWORD
    echo ""
fi

if [ -z "$NEW_PASSWORD" ]; then
    echo -e "${RED}❌ Error: Password cannot be empty${NC}"
    exit 1
fi

if [ ${#NEW_PASSWORD} -lt 6 ]; then
    echo -e "${RED}❌ Error: Password must be at least 6 characters${NC}"
    exit 1
fi

# Get role
if [ -n "$3" ]; then
    ROLE="$3"
else
    read -p "Enter role (user/admin) [user]: " ROLE
    ROLE="${ROLE:-user}"
fi

if [ "$ROLE" != "user" ] && [ "$ROLE" != "admin" ]; then
    echo -e "${RED}❌ Error: Role must be 'user' or 'admin'${NC}"
    exit 1
fi

# Check if user already exists
echo -e "${YELLOW}Checking if user exists...${NC}"
USER_EXISTS=$(docker run --rm --network host \
    postgres:16-alpine \
    psql "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" \
    -t -c "SELECT COUNT(*) FROM users WHERE username = '${USERNAME}';" 2>/dev/null | tr -d ' ')

if [ "$USER_EXISTS" != "0" ]; then
    echo -e "${RED}❌ Error: User '${USERNAME}' already exists${NC}"
    exit 1
fi

# Generate bcrypt hash using Python Docker image
echo -e "${YELLOW}Generating password hash...${NC}"
HASHED_PASSWORD=$(docker run --rm python:3.12-slim sh -c "
pip install -q bcrypt 2>/dev/null
python3 -c \"
import bcrypt
password = '${NEW_PASSWORD}'.encode('utf-8')
hashed = bcrypt.hashpw(password, bcrypt.gensalt(10))
print(hashed.decode('utf-8'))
\"
")

if [ -z "$HASHED_PASSWORD" ]; then
    echo -e "${RED}❌ Error: Failed to generate password hash${NC}"
    exit 1
fi

# Create user in database
echo -e "${YELLOW}Creating user...${NC}"
docker run --rm --network host \
    postgres:16-alpine \
    psql "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" \
    -c "INSERT INTO users (username, password_hash, role) VALUES ('${USERNAME}', '${HASHED_PASSWORD}', '${ROLE}');" 2>/dev/null

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "${GREEN}✅ User created successfully!${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "   Username: ${USERNAME}"
echo "   Role:     ${ROLE}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

