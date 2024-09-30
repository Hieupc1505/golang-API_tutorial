.PHONY: postgres createdb dropdb migrateup migratedown sqlc mockgen

# Định nghĩa run seerrver
run: 
	go run main.go

# Định nghĩa target cho PostgreSQL
postgres:
	docker run --name pg16_todo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

# Định nghĩa target để tạo database
createdb:
	docker exec -it pg16_todo createdb --username=root --owner=root todo_list

# Định nghĩa target để xóa database
dropdb:
	docker exec -it pg16_todo dropdb todo_list

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_list?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/todo_list?sslmode=disable" -verbose down

test:
	go test -v ./...

sqlc:
	sqlc generate

mockgen: 
	mockgen -package=mockdb -destination=db/mock/store.go --build_flags=--mod=mod github.com/hieupc09/simple_api/db/sqlc Store