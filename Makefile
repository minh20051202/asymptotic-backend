build:
	@go build -o bin/ticket-system

run: build
	@./bin/ticket-system

test:
	@go test -v ./...