.PHONY: help build run test clean docker-build docker-up docker-down

help:
	@echo "Available commands:"
	@echo "  build         Build the Go application"
	@echo "  run           Run the application locally"
	@echo "  test          Run tests"
	@echo "  clean         Clean build artifacts"
	@echo "  docker-build  Build Docker image"
	@echo "  docker-up     Start Docker Compose"
	@echo "  docker-down   Stop Docker Compose"

build:
	cd backend && go build -o bin/app ./cmd/main.go

run:
	cd backend && go run ./cmd/main.go

test:
	cd backend && go test -v -cover ./...

clean:
	cd backend && rm -rf bin/
	cd backend && go clean -testcache

docker-build:
	cd backend && docker build -t user-crud-app .

docker-up:
	cd backend && docker-compose up -d

docker-down:
	cd backend && docker-compose down

lint:
	cd backend && golangci-lint run

security:
	cd backend && gosec ./...