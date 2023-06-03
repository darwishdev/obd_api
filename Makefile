include app.env

sqlc:
	rm -rf pkg/sqlc/gen/*.sql.go && sqlc generate

run:
	fresh


proto_push:
	cd pkg/proto && buf push




docker_push:
	docker tag obd_api exploremelon/mln_carbrain:prod && docker push exploremelon/mln_carbrain:prod



deploy:
	docker push exploremelon/mln_carbrain

certs:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout certs/server.key -out certs/server.crt -days 365

proto:
	rm -rf pkg/pb/obd/v1/* && buf generate
createdb:
	docker exec -it postgres  createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres  dropdb --username=${DB_USER}   ${DB_NAME}  --force

refreshdb:
	make dropdb && make createdb && make migrateup

refreshdbtest:
	make refreshdb -f test.Makefile


refreshdbcloud:
	make refreshdb -f cloud.Makefile


mock:
	mockgen -package mockdb -destination pkg/sqlc/mock/store.go github.com/darwishdev/obd_api/pkg/sqlc/gen Store

migratenew:
	migrate create -ext sql -dir pkg/db/migration -seq ${name}

migrateup:
	migrate -path pkg/db/migration -database ${DB_SOURCE} -verbose up

debug_db:
	docker exec -it postgres psql -U ${DB_USER} -d ${BD_NAME}

test:
	go test -v -cover -short ./...

migratedown:
	migrate -path pckg/db/migration -database ${DB_SOURCE} -verbose down


evans:
	evans --host grpc://mln_api.exploremelon.com  -r repl