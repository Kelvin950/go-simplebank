postgres:
	docker run --name some-postgres -e POSTGRES_USER=root   -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it some-postgres  createdb --username=root --owner=root simple_bank

migrateup:
	migrate  -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate  -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate  -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
	
migratedown1:
	migrate  -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 2

dropdb:
	docker  exec -it some-postgres dropdb simple_bank

sqlc:
	sqlc generate

test:
	go test -v  -cover ./...

serer:
	go run main.go

mock:
	mockgen -package mockdb  -source  db/store.go -destination db/mock/mock.go

.PHONY: postgres createdb  dropdb migrateup migratedown migrateup1 migratedown1  sqlc test server mock

