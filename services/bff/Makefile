SERVICE = bff

.PHONY: run log
run:
	docker compose up -d --build $(SERVICE)
log:
	docker compose logs -f $(SERVICE)
build:
	docker build .

.PHONY: lint lint-fix
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...