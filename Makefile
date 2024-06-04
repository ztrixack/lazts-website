#!make

.PHONY: setup test test-it test-it-docker coverage dev dev-docker run

ifeq ($(OS),Windows_NT)
  uname_S := Windows
else
  uname_S := $(shell uname -s)
endif

setup:
	@echo "Setting up pre-commit..."
	./scripts/setup-pre-commit.sh
	@echo "Setting up air..."
	go install github.com/air-verse/air@latest
	@echo "Setting up env..."
	cp .env.template .env

test:
	@echo "Running tests..."
	go test -v ./...

test-it:
	@echo "Running integration tests..."
	go test -v -run "Test.*IT" -tags=integration ./...

test-it-docker:
	docker compose -f docker-compose.it.test.yaml down && \
	docker compose -f docker-compose.it.test.yaml up --build --force-recreate --abort-on-container-exit --exit-code-from it-tests && \
	docker compose -f docker-compose.it.test.yaml rm -s -f -v

coverage:
	go test -cover -coverprofile=report.out -v ./...
	go tool cover -html=report.out -o coverage.html

dev:
	@echo "Running the development server..."
ifeq ($(uname_S), Windows)
	air -c ./devops/.air.win.toml
else
	air -c ./devops/.air.toml
endif

dev-docker:
	@echo "Running the docker for development ..."
	docker compose -f docker-compose.dev.yaml down && \
	docker compose -f docker-compose.dev.yaml up --build --force-recreate --abort-on-container-exit && \
	docker compose -f docker-compose.dev.yaml rm -s -f -v



run:
	@echo "Running the server..."
	go run ./cmd/app/main.go
