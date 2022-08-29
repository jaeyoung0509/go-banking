DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine 

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go


.PHONY: createdb createdb dropdb migrateup migratedown test  sqlc server migrateup1 migratedown1
