include cloud.env

createdb:
	docker exec -it postgres   createdb --username=${DB_USER} --host=${DB_HOST}  --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres   dropdb --username=${DB_USER}  -h ${DB_HOST}  ${DB_NAME} --force


migrateup:
	migrate -path pkg/db/migration -database ${DB_SOURCE} -verbose up

debug_db:
	docker exec -it postgres psql -U ${DB_USER} -d ${BD_NAME}

refreshdb:
	make dropdb -f cloud.Makefile   && make createdb -f cloud.Makefile   && make migrateup -f cloud.Makefile  
test:
	go test -v -cover -short ./...

migratedown:
	migrate -path pckg/db/migration -database ${DB_SOURCE} -verbose down

