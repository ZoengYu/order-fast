ORDERFAST_CONTAINER := order_fast_db
ORDERFAST_DB := order_fast
ORDERFAST_USER := runner
postgres:
	docker run --name ${ORDERFAST_CONTAINER} -p 5432:5432 -e POSTGRES_USER=runner -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it ${ORDERFAST_CONTAINER} createdb --username=runner --owner=runner ${ORDERFAST_DB}

dropdb:
	docker exec -it ${ORDERFAST_CONTAINER} dropdb -U ${ORDERFAST_USER} ${ORDERFAST_DB}

migrateup:
	migrate -path db/migration -database "postgresql://${ORDERFAST_USER}:password@localhost:5432/${ORDERFAST_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${ORDERFAST_USER}:password@localhost:5432/${ORDERFAST_DB}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

tests:
	go test -v -cover ./tests

.PHOMY: postgres createdb dropdb migrateup migratedown sqlc
