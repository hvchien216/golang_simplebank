postgres:
	echo 'running container ...'
createdb:
	docker exec -it postgres psql -U postgres -c "create database simple_bank owner postgres;"
dropdb:
	docker exec -it postgres psql -U postgres -c "drop database simple_bank;"
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
