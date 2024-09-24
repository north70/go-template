SERVICE_NAME := go-template

GO_OUT_DIR := internal/pb
BIN_DIR := $(CURDIR)/bin
CONFIG_FILE := config.yaml
GOOGLEAPIS_DIR := $(BIN_DIR)/third_party/googleapis
GRPC_GATEWAY_DIR := $(BIN_DIR)/third_party/grpc-gateway

POSTGRES_DSN := "postgres://user:password@postgres:5432/go_template?sslmode=disable"

PROTO_CONFIG := configs/proto.yml

.PHONY: install-tools
install-tools:
	@mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BIN_DIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(BIN_DIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(BIN_DIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(BIN_DIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(BIN_DIR) go install github.com/vektra/mockery/v2@v2.46.0

.PHONY: get-dependencies
get-dependencies:
	@if [ ! -d "$(GOOGLEAPIS_DIR)" ]; then \
		echo "Cloning googleapis..."; \
		git clone --depth 1 https://github.com/googleapis/googleapis.git $(GOOGLEAPIS_DIR); \
	else \
		echo "googleapis already exists. Skipping clone."; \
	fi
	@if [ ! -d "$(GRPC_GATEWAY_DIR)" ]; then \
		echo "Cloning grpc-gateway..."; \
		git clone --depth 1 https://github.com/grpc-ecosystem/grpc-gateway.git $(GRPC_GATEWAY_DIR); \
	else \
		echo "grpc-gateway already exists. Skipping clone."; \
	fi

.PHONY: generate-proto
generate-proto: get-dependencies
	@echo "Reading proto directories from $(PROTO_CONFIG)..."
	@for dir in $$(yq e '.proto_dirs[]' $(PROTO_CONFIG)); do \
		OUT_DIR=$(GO_OUT_DIR)/$$dir; \
		mkdir -p $$OUT_DIR; \
		if [ "$$dir" = "$(SERVICE_NAME)" ]; then \
			OPENAPI_DIR=$(OPENAPI_OUT_DIR)/$$dir; \
			mkdir -p $$OPENAPI_DIR; \
			echo "Generating proto and OpenAPI files for $$dir..."; \
			protoc -I proto/$$dir -I $(GOOGLEAPIS_DIR) -I $(GRPC_GATEWAY_DIR) \
				--plugin=protoc-gen-go=$(BIN_DIR)/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=$(BIN_DIR)/protoc-gen-go-grpc \
				--plugin=protoc-gen-grpc-gateway=$(BIN_DIR)/protoc-gen-grpc-gateway \
				--plugin=protoc-gen-openapiv2=$(BIN_DIR)/protoc-gen-openapiv2 \
				--go_out=$$OUT_DIR --go_opt=paths=source_relative \
				--go-grpc_out=$$OUT_DIR --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=$$OUT_DIR --grpc-gateway_opt=paths=source_relative \
				--openapiv2_out=$(GO_OUT_DIR) \
				--openapiv2_opt=logtostderr=true \
				--openapiv2_opt=allow_merge=true \
				--openapiv2_opt=merge_file_name="$(SERVICE_NAME)" \
				proto/$$dir/*.proto; \
		else \
			echo "Generating proto files for $$dir..."; \
			protoc -I proto/$$dir -I $(GOOGLEAPIS_DIR) -I $(GRPC_GATEWAY_DIR) \
				--plugin=protoc-gen-go=$(BIN_DIR)/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=$(BIN_DIR)/protoc-gen-go-grpc \
				--plugin=protoc-gen-grpc-gateway=$(BIN_DIR)/protoc-gen-grpc-gateway \
				--go_out=$$OUT_DIR --go_opt=paths=source_relative \
				--go-grpc_out=$$OUT_DIR --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=$$OUT_DIR --grpc-gateway_opt=paths=source_relative \
				proto/$$dir/*.proto; \
		fi \
	done

.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	@$(BIN_DIR)/golangci-lint run ./... --config .golangci.yml

.PHONY: docker-up
start:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "Docker containers are up and running"

.PHONY: docker-down
stop:
	@echo "Stopping Docker containers..."
	docker-compose down
	@echo "Docker containers have been stopped"

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations up..."
	docker-compose build migrations
	docker-compose run --rm -e POSTGRES_DSN=$(POSTGRES_DSN) migrations goose -dir /app/migrations postgres "$(POSTGRES_DSN)" up

.PHONY: migrate-down
migrate-down:
	@echo "Running migrations down..."
	docker-compose run --rm -e POSTGRES_DSN=$(POSTGRES_DSN) migrations goose -dir /app/migrations postgres "$(POSTGRES_DSN)" down

.PHONY: migrate-status
migrate-status:
	@echo "Checking migration status..."
	docker-compose run --rm -e POSTGRES_DSN=$(POSTGRES_DSN) migrations goose -dir /app/migrations postgres "$(POSTGRES_DSN)" status

.PHONY: test
test:
	@echo "Running tests..."
	go test ./...