build:
	@go build -o bin/asymptotic-backend

run: build
	@./bin/asymptotic-backend

test:
	@go test -v ./...