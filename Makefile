.PHONY: run build test clean docker-up docker-down

DATABASE_URL=postgresql://postgres:postgres@localhost:5432/noveats_db?sslmode=disable

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

deps:
	go mod download
	go mod tidy

# === MIGRATION COMMANDS ===
# Run all pending migrations
migrate-up:
	@echo "Running migrations..."
	migrate -database "$(DATABASE_URL)" -path migrations up
	@echo "Migrations completed"

# Rollback last migration
migrate-down:
	@echo "Rolling back last migration..."
	migrate -database "$(DATABASE_URL)" -path migrations down 1
	@echo "Rollback completed"

# Rollback all migrations
migrate-down-all:
	@echo "Rolling back all migrations..."
	migrate -database "$(DATABASE_URL)" -path migrations down -all
	@echo "All migrations rolled back"

# Create a new migration file
# Usage: make migrate-create name=add_user_status
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: name is required. Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)
	@echo "Migration files created"

# Force migration version (use with caution)
# Usage: make migrate-force version=1
migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Error: version is required. Usage: make migrate-force version=1"; \
		exit 1; \
	fi
	migrate -database "$(DATABASE_URL)" -path migrations force $(version)

# Check current migration version
migrate-version:
	@migrate -database "$(DATABASE_URL)" -path migrations version

# Go to specific migration version
# Usage: make migrate-goto version=1
migrate-goto:
	@if [ -z "$(version)" ]; then \
		echo "Error: version is required. Usage: make migrate-goto version=1"; \
		exit 1; \
	fi
	migrate -database "$(DATABASE_URL)" -path migrations goto $(version)

# Check database status
db-status:
	@docker exec -it noveats_postgres psql -U postgres -d noveats_db -c "\dt"

# Access database shell
db-shell:
	@docker exec -it noveats_postgres psql -U postgres -d noveats_db

# Drop and recreate database (CAUTION: destroys all data)
db-reset:
	@echo "WARNING: This will destroy all data!"
	@echo "Press Ctrl+C to cancel, or Enter to continue..."
	@read confirm
	docker exec -it noveats_postgres psql -U postgres -c "DROP DATABASE IF EXISTS noveats_db;"
	docker exec -it noveats_postgres psql -U postgres -c "CREATE DATABASE noveats_db;"
	@echo "Database reset completed"
	@echo "Run 'make migrate-up' to create tables"
