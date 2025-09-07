postgres:
	podman run --name some-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

createdb:
	podman exec -it some-postgres createdb --username=root --owner=root simple_bank

dropdb:
	podman exec -it some-postgres dropdb simple_bank

migrateup:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" --verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
