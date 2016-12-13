BIN=snpt
BUILD_DIR=./build
SRC_DIR=./src

.PHONY: default
default: build

.PHONY: test
test:
	go test -v $(SRC_DIR)/...

.PHONY: lint
lint:
	gometalinter \
		--disable=errcheck \
		--disable=gotype \
		--enable=gofmt \
		--vendor \
		$(SRC_DIR)/...

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BIN) $(SRC_DIR)

.PHONY: build-all
build-all:
	gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" -os="darwin windows linux" -arch="amd64" $(SRC_DIR)/...

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
	glide install

.PHONY: install-env-deps
install-env-deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/Masterminds/glide
	gometalinter --install

.PHONY: fmt
fmt:
	go fmt $(SRC_DIR)/...
