run:
	go run main.go

test: 
	go test -v -cover ./...

cassandra:
	docker-compose -f docker-compose.db.yml up

cassandra-stop:
	docker-compose -f docker-compose.db.yml down

.PHONY: run test cassandra cassandra-stop
