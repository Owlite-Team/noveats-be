#!/bin/bash

# Create new migration helper script
set -e

if [ -z "$1" ]; then
    echo "Error: migration name required"
    echo "Usage: ./scripts/create-migration.sh MIGRATION_NAME"
    echo "Example: ./scripts/create-migration.sh add_user_status"
    exit 1
fi

migrate create -ext sql -dir migrations -seq "$1"
echo "✅ Migration files created for: $1"
ls -la migrations/ | tail -2
