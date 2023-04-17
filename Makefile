# Build
.PHONY: build
make build:
	rm -f main
	go build ./cmd/app/main.go
	./main

.PHONY: run
make run:
	./main

.PHONY: build-docker
make build-docker:
	docker compose up

# Linter
.PHONY: lint
make lint:
	golangci-lint run ./...

# Test
.PHONY: test
make test:
	go test -cover ./...