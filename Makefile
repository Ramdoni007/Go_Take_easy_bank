postgres: 
	docker run --name postgresqlx --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15-alpine

createdb: 
		docker exec -it postgresql createdb --username=root --owner=root easy_bank

dropdb:
		docker exec -it postgresql dropdb easy_bank

migrateup: 
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank" -verbose up

migrateup1:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable" -verbose up 1

migratedown: 
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable" -verbose down

migratedown1:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable" -verbose down 1

sqlc:
		sqlc generate

test:
		go test -v -cover ./...

server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/ramdoni007/Take_Easy_Bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server
