.PHONY: all

default: test lint

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

test:
	go test ./... -cover -race
