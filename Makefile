postgres:
	docker run --name postgres16.2 -p 5433:5432 -e POSTGRES_USER=anoop -e POSTGRES_PASSWORD=go -e POSTGRES_DB=golang -d postgres:16.2-alpine

createdb:
	docker exec -it postgres16.2 createdb -U anoop --owner anoop simple_bank

dropdb:
	docker exec -it postgres16.2 dropdb simple_bank

migrateup:
	migrate --path db/migration --database "postgresql://anoop:go@localhost:5433/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://anoop:go@localhost:5433/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, sqlc, test