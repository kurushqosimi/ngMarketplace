postgres:
	docker run --name postgres --network marketplace-network -p 5432:5432 -e POSTGRES_USER=marketplace -e POSTGRES_PASSWORD=marketplace -d postgres:17.2-alpine3.21

createdb:
	docker exec -it postgres createdb --username=marketplace --owner=marketplace marketplace_db

dropdb:
	docker exec -it postgres dropdb marketplace

migrateup:
	migrate -path db/migration -database "postgresql://marketplace:marketplace@localhost:5432/marketplace_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://marketplace:marketplace@localhost:5432/marketplace_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://marketplace:marketplace@localhost:5432/marketplace_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://marketplace:marketplace@localhost:5432/marketplace_db?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

server:
	go run main.go

docker-compose-restart:
	docker compose down
	docker rmi ngmarketplace-api
	docker compose up -d

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 server docker-compose-restart