.PHONY: infra infra-stop migrations-up migrations-down run swag test lint install-hooks

DB_NAME=effective_mobile
DB_USER=postgres
DB_PASS=postgres
DB_PORT=5432
MIGRATION_FOLDER=./migrations

infra:
	docker compose up -d --build --force-recreate --wait

infra-stop:
	docker compose down

migrations-up:
	goose postgres 'host=localhost port=${DB_PORT} user=${DB_USER} password=${DB_PASS} sslmode=disable dbname=${DB_NAME}' -dir ${MIGRATION_FOLDER} up

migrations-down:
	goose postgres 'host=localhost port=${DB_PORT} user=${DB_USER} password=${DB_PASS} sslmode=disable dbname=${DB_NAME}' -dir ${MIGRATION_FOLDER} down

run:
	go run ./cmd/app/main.go

swag:
	swag init -g cmd/app/main.go -o docs

test:
	go test ./... -race

lint:
	golangci-lint run ./...

install-hooks:
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
	@echo "pre-commit hook установлен"

