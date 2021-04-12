MIGRATIONS_FOLDER=internal/db/migrations

DATABASE_URL='postgres://postgres:1234567@localhost:5430/server-db?sslmode=disable'

compose-up:
	docker-compose up -d

docs-generate:
	swag init -g cmd/main.go

migration-up:
	migrate -database ${DATABASE_URL} -path ${MIGRATIONS_FOLDER} up

migration-down:
	migrate -database ${DATABASE_URL} -path ${MIGRATIONS_FOLDER} down

build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go