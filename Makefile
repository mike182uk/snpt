BIN=snpt
BUILD_DIR=./build
ENTRYPOINT=./cmd/snpt

.PHONY: default
default: build

.PHONY: test
test:
	GO111MODULE=on go test -v ./internal/...

.PHONY: coverage
coverage:
	GO111MODULE=on go test -v -covermode=count -coverprofile=coverage.out ./internal/...
	goveralls -coverprofile=coverage.out -service=travis-ci

.PHONY: lint
lint:
	GO111MODULE=on golangci-lint run

.PHONY: build
build:
	GO111MODULE=on go build -o $(BUILD_DIR)/$(BIN) $(ENTRYPOINT)

.PHONY: build-all
build-all:
	GO111MODULE=on gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" -os="darwin windows linux" -arch="amd64" $(ENTRYPOINT)

.PHONY: package
package:
	find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: install
install: install-tools
	GO111MODULE=on go mod download

.PHONY: install-tools
install-tools:
	GO111MODULE=off \
	go get -u github.com/mitchellh/gox \
		github.com/mattn/goveralls \
		github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt:
	go fmt ./internal/... ./cmd/...
