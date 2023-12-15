SERVICES := bff item mission tree user

.PHONY: setup setup/env setup/lint setup/work setup/wire setup/protoc setup/creds
setup: setup/go setup/env setup/lint setup/work setup/wire setup/protoc setup/creds
	@echo "Setup completed"
# go version check 1.21.5
setup/go:
	@if go version | grep -q "go1.21.5"; then \
		echo "Go 1.21.5 is already installed."; \
	else \
		echo "Go 1.21.5 is not installed. Please install go1.21.5"; \
		exit 1; \
	fi
# .env file init
setup/env:
	cp .env.sample .env
# golangci-lint install
setup/lint: setup/go
	@if ! command -v golangci-lint &>/dev/null; then \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed."; \
	fi
# work file init
setup/work: setup/go
	@if [ ! -f go.work ]; then \
		go work init; \
	fi
	go work use $(SERVICES)
# wire(DI tool) install
setup/wire: setup/go
	go install github.com/google/wire/cmd/wire@latest
# protobuf-compiler & go plugin install
setup/protoc: setup/go
	@if ! command -v protoc &>/dev/null; then \
		echo "protoc is not installed. Please, Install from https://grpc.io/docs/protoc-installation/" \
		exit 1; \
	else \
		echo "protoc is already installed."; \
	fi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest;
# check service-account.json
setup/creds:
	@if [ ! -f ./service-account.json ]; then \
		echo "service-account.json is not exist. Please, download from GCP"; \
		exit 1; \
	fi

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
		echo "===================="; \
		echo "Running lint for $$service"; \
		echo "===================="; \
		golangci-lint run ./$$service/...; \
	done
lint-fix:
	@for service in $(SERVICES); do \
		echo "===================="; \
		echo "Running lint for $$service"; \
		echo "===================="; \
		golangci-lint run --fix ./$$service/...; \
	done

# Wire
.PHONY: wire wire-gen
wire:
	@for service in $(SERVICES); do \
		echo "Running wire for $$service"; \
		wire ./$$service/...; \
	done

# Go Mod
.PHONY: tidy 
tidy:
	@for service in $(SERVICES); do \
		cd $$service; \
		go mod tidy; \
		cd ..; \
	done