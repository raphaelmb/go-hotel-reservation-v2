build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

.PHONY: build run