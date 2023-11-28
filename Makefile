include .env

MIGRATION_PATH = "./migrations"
MIGRATION_NAME = "migration"

all: start

start: install-dependences compose-up migrate-up	
	go install
	notesServer

compose-up:
	docker compose up -d || true
	@sleep 1

migrate: migrate-up
	
migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) down

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(MIGRATION_NAME) 

install-dependences: 
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest	
