DB_URL=postgres://postgres:postgres@localhost:5433/supportportal_development?sslmode=disable
APP_BIN=bin/supportportal

.PHONY: run build test vet fmt fmt-check lint ci compose-up compose-down db-up db-down migrate-up migrate-down migrate-status

run:
	go run ./cmd/api

build:
	go build -o $(APP_BIN) ./cmd/api

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)"

lint:
	golangci-lint run ./...

ci: fmt-check vet test build

compose-up:
	docker compose build --no-cache && docker compose up -d

compose-down:
	docker compose down

db-up:
	docker compose up -d postgres

db-down:
	docker compose stop postgres

migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down

migrate-status:
	goose -dir migrations postgres "$(DB_URL)" status
