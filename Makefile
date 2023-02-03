postgres:
	echo 'running container ...'
createdb:
	docker exec -it postgres psql -U postgres -c "create database simple_bank owner postgres;"
dropdb:
	docker exec -it postgres psql -U postgres -c "drop database simple_bank;"
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	 mockgen -package mockdb -destination db/mock/store.go simple_bank/db/sqlc Store
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto
evans:
	 evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 proto
