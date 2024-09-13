include .env
.PHONY: build run test create-swagger compose-up compose-down

build:
	go build -o ./bin/main ./cmd/main.go

run: build
	./bin/main

test:
	go test -v ./...

test-file:
	find . -name "*_test.go" -exec dirname {} \; | sort -u | xargs -I {} go test -v {}

create-swagger:
	swag init -g ./internal/adapter/http/http.go -o ./internal/adapter/http/docs --parseDependency

compose-up:
	rm -rf tmp && rm -rf ./cmd/tmp && docker-compose up -d

compose-down:
	rm -rf tmp && rm -rf ./cmd/tmp && docker-compose down

compose-build:
	rm -rf tmp && rm -rf ./cmd/tmp && docker-compose build

compose-rebuild:
	rm -rf tmp && rm -rf ./cmd/tmp && docker-compose build --no-cache

local:
	rm -rf tmp && rm -rf ./cmd/tmp && docker-compose --env-file .env up -d && docker-compose logs -f app