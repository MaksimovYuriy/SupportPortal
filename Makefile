DB_URL=postgres://postgres:postgres@localhost:5433/supportportal_development?sslmode=disable

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down

migrate-status:
	goose -dir migrations postgres "$(DB_URL)" status