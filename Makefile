.PHONY: migration-up migration-down api

migration-up:
	goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" up

migration-down:
	goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable" down

db-up:
	docker compose up -d

db-down:
	docker compose down

api:
	cd cmd && go run main.go

test:
	go test -v ./internal/...