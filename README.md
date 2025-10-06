# 🚀 How to Run the Project
### 1. Start PostgreSQL
make docker-up

### 2. Run migrations (creates tables)
make migrate-up

### 3. Start the application
make run

### 4. Test the API
curl http://localhost:8080/health

---

# 🗄️ Database Migrations
This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Run all pending migrations
make migrate-up

### Rollback last migration
make migrate-down

### Rollback all migrations
make migrate-down-all

### Check current migration version
make migrate-version

### Create new migration
make migrate-create name=add_user_status

### Force set migration version (if stuck)
make migrate-force version=1

### Go to specific version
make migrate-goto version=1

### Check database tables
make db-status

### Access database shell
make db-shell

### Reset database (⚠️ DESTROYS ALL DATA)
make db-reset

### Check without entering psql
docker exec -it noveats_postgres psql -U postgres -d noveats_db -c "\dt"
