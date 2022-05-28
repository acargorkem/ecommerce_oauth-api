run:
	go run main.go

run-docker:
	docker-compose up

build-docker:
	docker-compose up --build

stop-docker:
	docker-compose down

test: 
	go test -v -cover ./...

cassandra:
	docker-compose -f docker-compose.db.yml up

cassandra-stop:
	docker-compose -f docker-compose.db.yml down	

.PHONY: run run-docker build-docker stop-docker test cassandra cassandra-stop
