postgres: 
	docker run --name postgresql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15-alpine

createdb: 
		docker exec -it postgresql createdb --username=root --owner=root easy_bank

dropdb:
	docker exec -it postgresql dropdb easy_bank 

migrateup: 
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable" -verbose up

migratedown: 
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/easy_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ramdoni007/Take_Easy_Bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
