BIN=snpt
BUILD_DIR=./build
ENTRYPOINT=./cmd/snpt

.PHONY: default
default: build

.PHONY: test
test:
	go test -cover -v ./internal/...

.PHONY: coverage
coverage:
	goveralls -package=./internal/...

.PHONY: lint
lint:
	gometalinter \
		--enable=gofmt \
		--enable=misspell \
		--vendor \
		--deadline=180s \
		--exclude=internal/platform/storage/test.go \
		./internal/... ./cmd/...

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BIN) $(ENTRYPOINT)

.PHONY: build-all
build-all:
	gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" -os="darwin windows linux" -arch="amd64" $(ENTRYPOINT)

.PHONY: package
package:
	find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: clean-all
clean-all: clean
	rm -rf ./vendor

.PHONY: install
install:
	dep ensure

.PHONY: install-env-deps
install-env-deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/mattn/goveralls
	gometalinter --install

.PHONY: fmt
fmt:
	go fmt ./internal/... ./cmd/...
