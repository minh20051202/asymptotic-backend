build:
	@go build -o bin/asymptotic-backend ./cmd/api

run: build
	@./bin/asymptotic-backend

test:
	@go test -v ./...