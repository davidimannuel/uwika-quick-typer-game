#!/bin/bash

# Script to create default admin user
# Usage: ./create_admin.sh

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-s3cret}
DB_NAME=${DB_NAME:-quick_typer}

# Generate UUID
ADMIN_ID=$(uuidgen | tr '[:upper:]' '[:lower:]')

# Hash password (admin123)
# You need to install htpasswd or use Go to generate this
PASSWORD_HASH='$2a$10$N9qo8uLOickgx2ZMRZoMye1OzSTxkLpwDWAM.iFUfZaOX/LXPqWwC'

# SQL command
SQL="INSERT INTO users (user_id, username, password_hash, role) 
     VALUES ('$ADMIN_ID', 'admin', '$PASSWORD_HASH', 'admin') 
     ON CONFLICT (username) DO UPDATE SET password_hash = '$PASSWORD_HASH';"

# Execute SQL
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "$SQL"

echo "Admin user created/updated successfully!"
echo "Username: admin"
echo "Password: admin123"

