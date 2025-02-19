IS_WINDOWS=false
OUTPUT_FILE=go-kcloutie
DOCKER_EXE=podman
BUILD_DATE=NA
BUILD_COMMIT=$(shell git rev-parse HEAD)
BUILD_VERSION=$(shell git describe --abbrev=0 --tags)
BUILD_DATE="$$(date --iso=seconds)"
SERVER_CONFIG_FILE=tests/files/serverConfig.json
ifeq ($(OS),Windows_NT)
	IS_WINDOWS=true
	OUTPUT_FILE=go-kcloutie.exe

else



endif
LDFLAGS += -ldflags "-X github.com/IaC/go-kcloutie/pkg/params/version.BuildTime=$(BUILD_DATE) -X github.com/IaC/go-kcloutie/pkg/params/version.BuildVersion=$(BUILD_VERSION) -X github.com/IaC/go-kcloutie/pkg/params/version.Commit=$(BUILD_COMMIT)"

.PHONY: build
build:
	@echo ""
	@echo "Is Windows: ${IS_WINDOWS}"
	@echo "BUILD_DATE: ${BUILD_DATE}"
	@echo "BUILD_COMMIT: ${BUILD_COMMIT}"
	@echo "BUILD_VERSION: ${BUILD_VERSION}"
	@echo "OUTPUT_FILE: ${OUTPUT_FILE}"
	@echo ""
	@echo "Building CLI..."
	@go build $(LDFLAGS) -o bin/$(OUTPUT_FILE) cmd/gokcloutie/gokcloutie.go

.PHONY: api-server
api-server: build
	@echo "Running Server..."
	@./$(OUTPUT_FILE) run server -c $(SERVER_CONFIG_FILE)

.PHONY: unit-test
unit-test: build
	@echo "Running Unit Tests..."
	@go test ./...

.PHONY: release
release: unit-test
	@echo "Releasing Product..."
	@echo "TAG: ${TAG}"
	@git tag -a $(Tag) -m "Release ${TAG}"
	@goreleaser --rm-dist

.PHONY: docs
docs:
	@echo "Generating Docs..."
	@go run ./cmd/gen-docs --standard --doc-path docs/go-kcloutie
	@swag init --generalInfo ./cmd/gokcloutie/gokcloutie.go

.PHONY: docs-custom
docs-custom:
	@echo "Generating Custom Docs and Swagger files..."
	@go run ./cmd/gen-docs --custom --doc-path docs/go-kcloutie-custom
	@swag init --generalInfo ./cmd/gokcloutie/gokcloutie.go

.PHONY: build-image
build-image:
	@echo "Building Container Image..."
	@$(DOCKER_EXE) build -t go-kcloutie:$(BUILD_VERSION) .
	@echo "$(DOCKER_EXE) run go-kcloutie:$(BUILD_VERSION) /opt/app-root/go-kcloutie version"
