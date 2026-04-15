# go-ecommerce-backend-api/Makefile
APP_NAME=server

dev:
	air

run:
	go run ./cmd/server/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/server

start: build
	./bin/$(APP_NAME)

tidy:
	go mod tidy

clean:
	rm -rf bin

.PHONY: run build start tidy clean dev