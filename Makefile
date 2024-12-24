DB_NAME=tempi
DB_USER=postgres
DB_PASSWORD=faxa
DB_HOST=localhost
DB_PORT=5432

run: build
	@./bin/billy-flow.exe

tidy:
	@go mod tidy

build:
	go mod tidy
	@echo "Make: Forming database..."
	@export PGPASSWORD=$(DB_PASSWORD) && psql -h $(DB_HOST) -U $(DB_USER) -p $(DB_PORT) -d postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);"
	@export PGPASSWORD=$(DB_PASSWORD) && psql -h $(DB_HOST) -U $(DB_USER) -p $(DB_PORT) -d postgres -c "CREATE DATABASE $(DB_NAME);"
	go build -o bin/billy-flow.exe cmd/main.go

clean:
	@rm -rf bin