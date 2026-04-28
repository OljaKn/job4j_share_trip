GO := go
PKG := ./...

.PHONY: deps fmt test lint build run up down migrate-up migrate-down migrate-status check

deps:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
fmt:
	go fmt ./...
test:
	go test ./...
lint:
	golangci-lint run
build:
	go build -o bin/sharetrip ./cmd/sharetrip
run:
	go run ./cmd/sharetrip
up:
	docker-compose up
down:
	docker-compose down
migrate-up:
	goose -dir migrations postgres "postgres://postgres:password@localhost:6543/sharetrip?sslmode=disable" up
migrate-down:
	goose -dir migrations postgres "postgres://postgres:password@localhost:6543/sharetrip?sslmode=disable" down
migrate-status:
	goose -dir migrations postgres "postgres://postgres:password@localhost:6543/sharetrip?sslmode=disable" status
check:
	fmt lint test
e2e:
	curl http://localhost:8080/ready