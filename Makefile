run:
	go run main.go

test: 
	go test -v -cover ./...

.PHONY: run test
