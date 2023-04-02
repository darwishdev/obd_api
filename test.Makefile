include test.env

createdb:
	docker exec -it postgres  createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres  dropdb --username=${DB_USER}   ${DB_NAME}


migrateup:
	migrate -path pkg/db/migration -database ${DB_SOURCE} -verbose up

debug_db:
	docker exec -it postgres psql -U ${DB_USER} -d ${BD_NAME}

refreshdb:
	make dropdb -f test.Makefile   && make createdb -f test.Makefile   && make migrateup -f test.Makefile  
test:
	go test -v -cover -short ./...

migratedown:
	migrate -path pckg/db/migration -database ${DB_SOURCE} -verbose down

