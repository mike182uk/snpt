BIN=snpt
BUILD_DIR=./build
ENTRYPOINT=./cmd/snpt

.PHONY: test
test: ## Run the tests
	GO111MODULE=on go test -v ./internal/...

.PHONY: coverage
coverage: ## Generate test coverage
	GO111MODULE=on go test -v -covermode=count -coverprofile=coverage.out ./internal/...
	goveralls -coverprofile=coverage.out -service=travis-ci

.PHONY: lint
lint: ## Lint the soruce files
	GO111MODULE=on golangci-lint run

.PHONY: build
build: ## Build the project for the current architecture
	GO111MODULE=on go build -o $(BUILD_DIR)/$(BIN) $(ENTRYPOINT)

.PHONY: build-all
build-all: ## Build the project for all supported architectures
	GO111MODULE=on gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" -os="darwin windows linux" -arch="amd64" $(ENTRYPOINT)

.PHONY: package
package: ## Package the any built binaries ready for distribution
	find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

.PHONY: clean
clean: ## Clean the workspace
	rm -rf $(BUILD_DIR)

.PHONY: install
install: install-tools ## Install project dependencies (including any required tools)
	GO111MODULE=on go mod download

.PHONY: install-tools
install-tools: ## Install tools required by the project
	GO111MODULE=off \
	go get -u github.com/mitchellh/gox \
		github.com/mattn/goveralls \
		github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt: ## Format the soruce files
	go fmt ./internal/... ./cmd/...

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
