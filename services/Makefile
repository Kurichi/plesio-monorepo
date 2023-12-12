SERVICES := bff item mission tree user

.PHONY: setup setup/env setup/lint
setup: setup/env setup/lint
setup/env:
	@echo "Setting up .env file with sample values"
	cp .env.sample .env
setup/lint:
	@if ! command -v golangci-lint &>/dev/null; then \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed."; \
	fi
setup/work:
	@if ! find . -name go.work -type f | grep -q go.work; then \
		go work init; \
	fi
	go work use $(SERVICES)

# Docker run
.PHONY: run log 
run: 
	docker compose up -d --build
log:
	docker compose logs -f

# Lint
.PHONY: lint lint-fix
lint: 
	@for service in $(SERVICES); do \
		golangci-lint run ./$$service/...; \
	done
lint-fix:
	@for service in $(SERVICES); do \
		golangci-lint run --fix ./$$service/...; \
	done

# Wire
.PHONY: wire wire-gen
wire:
	@for service in $(SERVICES); do \
		echo "Running wire for $$service"; \
		wire ./$$service/...; \
	done

sync:
	@for service in $(SERVICES); do \
		cd $$service; \
		go mod tidy; \
		cd ..; \
	done