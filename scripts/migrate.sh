#!/bin/bash

# Migration helper script
set -e

# Check if we're running in Docker or locally
if [ "$DOCKER_ENV" = "true" ]; then
    DATABASE_URL="postgresql://postgres:postgres@postgres:5432/noveats_db?sslmode=disable"
else
    DATABASE_URL="postgresql://postgres:postgres@127.0.0.1:5432/noveats_db?sslmode=disable"
fi

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
    docker)
        echo "Running migrations via Docker..."
        docker run --rm --network noveats-be_noveats_network \
            -v "$(pwd)/migrations:/migrations" \
            migrate/migrate:v4.16.2 \
            -database "postgresql://postgres:postgres@postgres:5432/noveats_db?sslmode=disable" \
            -path /migrations up
        echo "✅ Migrations completed via Docker"
        ;;
    *)
        echo "Usage: ./scripts/migrate.sh {up|down|version|force VERSION|docker}"
        echo "  up     - Run migrations (local)"
        echo "  down   - Rollback last migration (local)"
        echo "  version - Show current migration version"
        echo "  force  - Force migration to specific version"
        echo "  docker - Run migrations via Docker (recommended when DB is in Docker)"
        exit 1
        ;;
esac
