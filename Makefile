.PHONY: run build test clean docker-up docker-down migrate-up-docker migrate-down-docker migrate-down-all-docker migrate-up migrate-down migrate-create migrate-force migrate-version help

DATABASE_URL=postgresql://postgres:postgres@localhost:5432/noveats_db?sslmode=disable
POSTGRES_CONTAINER=noveats_postgres
DB_NAME=noveats_db

run:
	@echo "🚀 Starting application..."
	go run cmd/server/main.go

build:
	@echo "🔨 Building application..."
	go build -o bin/server cmd/server/main.go
	@echo "✅ Build complete"

test:
	@echo "🧪 Running tests..."
	go test -v ./...

test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -v -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean:
	@echo "🧹 Cleaning..."
	rm -rf bin/
	rm -f coverage.out
	@echo "✅ Clean complete"

deps:
	@echo "📦 Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies updated"

fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

lint:
	@echo "🔍 Running linter..."
	golangci-lint run ./...

vet:
	@echo "🔍 Running go vet..."
	go vet ./...

# === DOCKER COMMANDS ===
docker-up:
	@echo "🐳 Starting PostgreSQL..."
	docker-compose up -d postgres
	@echo "⏳ Waiting for database to be ready..."
	@sleep 3
	@echo "✅ PostgreSQL is running"

docker-down:
	@echo "🛑 Stopping Docker services..."
	docker-compose down
	@echo "✅ Docker services stopped"

docker-up-all:
	@echo "🐳 Starting all services..."
	docker-compose up -d
	@echo "✅ All services are running"

docker-logs:
	docker-compose logs -f

docker-logs-app:
	docker logs -f $(APP_CONTAINER)

docker-logs-db:
	docker logs -f $(POSTGRES_CONTAINER)

docker-restart:
	@echo "🔄 Restarting services..."
	docker-compose restart
	@echo "✅ Services restarted"

docker-build:
	@echo "🔨 Building Docker image..."
	docker-compose build
	@echo "✅ Docker image built"

migrate-up-docker:
	@echo "📤 Running migrations via Docker..."
	@docker run --rm --network noveats-be_noveats_network \
		-v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.16.2 \
		-database "postgresql://postgres:postgres@postgres:5432/noveats_db?sslmode=disable" \
		-path /migrations up
	@echo "✅ Migrations completed successfully via Docker"

migrate-down-docker:
	@echo "📥 Rolling back last migration via Docker..."
	@docker run --rm --network noveats-be_noveats_network \
		-v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.16.2 \
		-database "postgresql://postgres:postgres@postgres:5432/noveats_db?sslmode=disable" \
		-path /migrations down 1
	@echo "✅ Rollback completed via Docker"

migrate-down-all-docker:
	@echo "⚠️  WARNING: This will rollback ALL migrations!"
	@echo "Press Ctrl+C to cancel, or Enter to continue..."
	@read confirm
	@docker run --rm --network noveats-be_noveats_network \
		-v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.16.2 \
		-database "postgresql://postgres:postgres@postgres:5432/noveats_db?sslmode=disable" \
		-path /migrations down -all
	@echo "✅ Rollback completed via Docker"

# === MIGRATION COMMANDS ===
migrate-up:
	@echo "📤 Running migrations..."
	@migrate -database "$(DATABASE_URL)" -path migrations up
	@echo "✅ Migrations completed successfully"

migrate-down:
	@echo "📥 Rolling back last migration..."
	@migrate -database "$(DATABASE_URL)" -path migrations down 1
	@echo "✅ Rollback completed"

migrate-down-all:
	@echo "⚠️  WARNING: This will rollback ALL migrations!"
	@echo "Press Ctrl+C to cancel, or Enter to continue..."
	@read confirm
	@migrate -database "$(DATABASE_URL)" -path migrations down -all
	@echo "✅ All migrations rolled back"

# Create a new migration file
# Usage: make migrate-create name=add_user_status
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: name is required"; \
		echo "Usage: make migrate-create name=your_migration_name"; \
		echo "Example: make migrate-create name=create_products_table"; \
		exit 1; \
	fi
	@echo "📝 Creating migration: $(name)"
	@migrate create -ext sql -dir migrations -seq $(name)
	@echo "✅ Migration files created:"
	@ls -la migrations/ | tail -2

# Force migration version (use with caution)
# Usage: make migrate-force version=1
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "❌ Error: version is required"; \
		echo "Usage: make migrate-force version=VERSION_NUMBER"; \
		echo "Example: make migrate-force version=1"; \
		exit 1; \
	fi
	@echo "⚠️  Forcing migration to version $(version)..."
	@migrate -database "$(DATABASE_URL)" -path migrations force $(version)
	@echo "✅ Migration version set to $(version)"

migrate-version:
	@echo "📊 Current migration version:"
	@migrate -database "$(DATABASE_URL)" -path migrations version

# Go to specific migration version
# Usage: make migrate-goto version=1
migrate-goto:
	@if [ -z "$(version)" ]; then \
		echo "❌ Error: version is required"; \
		echo "Usage: make migrate-goto version=VERSION_NUMBER"; \
		echo "Example: make migrate-goto version=2"; \
		exit 1; \
	fi
	@echo "🎯 Migrating to version $(version)..."
	@migrate -database "$(DATABASE_URL)" -path migrations goto $(version)
	@echo "✅ Migrated to version $(version)"

## migrate-drop: Drop everything (⚠️ EXTREME CAUTION - DESTROYS ALL DATA)
migrate-drop:
	@echo "🔥 WARNING: This will DROP ALL TABLES and DATA!"
	@echo "Type 'yes' to confirm: "
	@read confirm; \
	if [ "$$confirm" = "yes" ]; then \
		migrate -database "$(DATABASE_URL)" -path migrations drop -f; \
		echo "✅ Database dropped"; \
	else \
		echo "❌ Cancelled"; \
	fi

# === DATABASE OPERATIONS COMMANDS ===
db-status:
	@echo "📊 Database Status:"
	@echo "\n📋 Tables:"
	@docker exec -it $(POSTGRES_CONTAINER) psql -U postgres -d $(DB_NAME) -c "\dt" || echo "⚠️  Database not accessible"
	@echo "\n🔄 Migration Version:"
	@make migrate-version 2>/dev/null || echo "⚠️  Cannot check migration version"

db-shell:
	@echo "🐘 Connecting to PostgreSQL..."
	@docker exec -it $(POSTGRES_CONTAINER) psql -U postgres -d $(DB_NAME)

## db-reset: Drop and recreate database (⚠️ DESTROYS ALL DATA)
db-reset:
	@echo "🔥 WARNING: This will DESTROY ALL DATA in the database!"
	@echo "Type 'yes' to confirm: "
	@read confirm; \
	if [ "$$confirm" = "yes" ]; then \
		echo "🗑️  Dropping database..."; \
		docker exec -it $(POSTGRES_CONTAINER) psql -U postgres -c "DROP DATABASE IF EXISTS noveats_db;"; \
		echo "📦 Creating database..."; \
		docker exec -it $(POSTGRES_CONTAINER) psql -U postgres -c "CREATE DATABASE noveats_db;"; \
		echo "✅ Database reset complete"; \
		echo "Run 'make migrate-up' to create tables"; \
	else \
		echo "❌ Cancelled"; \
	fi

## db-backup: Backup database to file
db-backup:
	@echo "💾 Creating database backup..."
	@mkdir -p backups
	@docker exec -t $(POSTGRES_CONTAINER) pg_dump -U postgres $(DB_NAME) > backups/backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "✅ Backup created in backups/ directory"

## db-restore: Restore database from latest backup
db-restore:
	@echo "📥 Restoring from latest backup..."
	@if [ -z "$$(ls -t backups/*.sql 2>/dev/null | head -1)" ]; then \
		echo "❌ No backup files found in backups/"; \
		exit 1; \
	fi
	@LATEST=$$(ls -t backups/*.sql | head -1); \
	echo "Restoring from: $$LATEST"; \
	docker exec -i $(POSTGRES_CONTAINER) psql -U postgres $(DB_NAME) < $$LATEST
	@echo "✅ Database restored"

## db-seed: Seed database with sample data (if seed file exists)
db-seed:
	@if [ -f "scripts/seed.sql" ]; then \
		echo "🌱 Seeding database..."; \
		docker exec -i $(POSTGRES_CONTAINER) psql -U postgres $(DB_NAME) < scripts/seed.sql; \
		echo "✅ Database seeded"; \
	else \
		echo "⚠️  No seed file found at scripts/seed.sql"; \
	fi


# === FULL WORKFLOW ===
## setup: Complete project setup (install deps, start db, run migrations)
setup:
	@echo "🚀 Setting up project..."
	@make deps
	@make docker-up
	@echo "⏳ Waiting for database..."
	@sleep 5
	@make migrate-up
	@echo "✅ Setup complete! Run 'make run' to start the app"

## dev: Start development environment (db + migrations + app)
dev:
	@echo "🚀 Starting development environment..."
	@make docker-up
	@sleep 3
	@make migrate-up
	@make run

## fresh: Fresh start (reset db, run migrations, start app)
fresh:
	@echo "🔄 Fresh start..."
	@make db-reset
	@sleep 2
	@make migrate-up
	@make run

## reset-dev: Reset development environment (stop, reset db, start)
reset-dev:
	@echo "🔄 Resetting development environment..."
	@make docker-down
	@make docker-up
	@sleep 3
	@make db-reset
	@make migrate-up
	@echo "✅ Development environment reset"

# === HELP ===
help:
	@echo "📚 Available commands:"
	@echo ""
	@echo "Development:"
	@grep -E '^## [a-zA-Z_-]+:.*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; /^## [a-zA-Z_-]+:/ {printf "  \033[36m%-20s\033[0m %s\n", substr($$1, 4), $$2}'

.DEFAULT_GOAL := help
