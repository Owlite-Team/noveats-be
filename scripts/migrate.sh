#!/bin/bash

# Migration helper script
set -e

DATABASE_URL="postgresql://postgres:postgres@localhost:5432/myapp_db?sslmode=disable"

case "$1" in
    up)
        echo "Running migrations..."
        migrate -database "$DATABASE_URL" -path migrations up
        echo "✅ Migrations completed"
        ;;
    down)
        echo "Rolling back last migration..."
        migrate -database "$DATABASE_URL" -path migrations down 1
        echo "✅ Rollback completed"
        ;;
    version)
        migrate -database "$DATABASE_URL" -path migrations version
        ;;
    force)
        if [ -z "$2" ]; then
            echo "Error: version number required"
            echo "Usage: ./scripts/migrate.sh force VERSION"
            exit 1
        fi
        migrate -database "$DATABASE_URL" -path migrations force "$2"
        ;;
    *)
        echo "Usage: ./scripts/migrate.sh {up|down|version|force VERSION}"
        exit 1
        ;;
esac
