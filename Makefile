postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:dMVm4HwrMzpVtGPFdLLC@database-aurora-postgres.cluster-c1boodh9e4gj.us-east-1.rds.amazonaws.com:5432/simple_bank" --verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:dMVm4HwrMzpVtGPFdLLC@database-aurora-postgres.cluster-c1boodh9e4gj.us-east-1.rds.amazonaws.com:5432/simple_bank" --verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:dMVm4HwrMzpVtGPFdLLC@database-aurora-postgres.cluster-c1boodh9e4gj.us-east-1.rds.amazonaws.com:5432/simple_bank" --verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:dMVm4HwrMzpVtGPFdLLC@database-aurora-postgres.cluster-c1boodh9e4gj.us-east-1.rds.amazonaws.com:5432/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/Manuel11713/simple-bank/db/sqlc Store 

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server mock