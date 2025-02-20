linter:
	golangci-lint run

create-migration:
	migrate create -ext sql -dir schema/migration -seq base

migrateup:
	migrate -path schema/migration -database "postgres://fara@localhost:5432/sslmode=disable" -verbose up

migratedown:
	migrate -path schema/migration -database "postgres://fara@localhost:5432/sslmode=disable" -verbose down

migratefix:
	migrate -path database/migration/ -database "postgres://fara@localhost:5432/sslmode=disable" force 000001

.PHONY: create-migration migrateup migratedown migration_fix linter