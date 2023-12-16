SERVICE = user

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

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=import \
		--go-grpc_out=. --go-grpc_opt=paths=import \
		./schema/*.proto

.PHONY: migrate
migrate:
	docker compose exec ${SERVICE} go run migration/migration.go
