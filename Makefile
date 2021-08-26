BIN=snpt
BUILD_DIR=./build
ENTRYPOINT=./cmd/snpt

.PHONY: test
test: ## Run the tests
	go test -v ./internal/...

.PHONY: coverage
coverage: ## Generate test coverage
	go test -race -covermode atomic -coverprofile=coverage.out ./internal/...

.PHONY: lint
lint: ## Lint the soruce files
	golangci-lint run

.PHONY: proto
proto: ## Compile protocol buffers
	protoc --go_out=. internal/pb/snippet.proto

.PHONY: mocks
mocks: ## Generate mocks
	mockery --name=BucketKeyValueStore --recursive=true

.PHONY: build
build: ## Build the project for the current architecture
	go build -o $(BUILD_DIR)/$(BIN) $(ENTRYPOINT)

.PHONY: build-all
build-all: ## Build the project for all supported architectures
	gox -output "$(BUILD_DIR)/$(BIN)-$(TAG)-{{.OS}}-{{.Arch}}/$(BIN)" -os="darwin windows linux" -arch="amd64" $(ENTRYPOINT)

.PHONY: package
package: ## Package the any built binaries ready for distribution
	find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

.PHONY: clean
clean: ## Clean the workspace
	rm -rf $(BUILD_DIR) coverage.out

.PHONY: install
install: install-tools ## Install project dependencies (including any required tools)
	go mod download

.PHONY: install-tools
install-tools: ## Install tools required by the project
	go get -u github.com/mitchellh/gox github.com/vektra/mockery/.../ google.golang.org/protobuf/cmd/protoc-gen-go
	if [ -z "$(CI)" ]; then curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.42.0; fi

.PHONY: fmt
fmt: ## Format the soruce files
	go fmt ./internal/... ./cmd/...

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
