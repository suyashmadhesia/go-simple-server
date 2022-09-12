postgres:
	docker run --name postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:12-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down
migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up
sqlc:
	sqlc generate
test:
	go test -v ./db/sqlc
.PHONY: createdb dropdb migratedown migrateup postgres sqlc test