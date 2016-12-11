.PHONY: default test lint build build-all package clean clean-all install install-env-deps fmt

BIN=snpt
BUILD_DIR=./build
SRC_DIR=./src

default: build

test:
	go test -v $(SRC_DIR)/...

lint:
	gometalinter \
		--disable=errcheck \
		--disable=gotype \
		--vendor \
		$(SRC_DIR)/...

build:
	go build -o $(BUILD_DIR)/$(BIN) $(SRC_DIR)

build-all:
	gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" $(SRC_DIR)/...

package:
	find $(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

clean:
	rm -rf $(BUILD_DIR)

clean-all: clean
	rm -rf ./vendor

install:
	glide install

install-env-deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/Masterminds/glide
	gometalinter --install

fmt:
	go fmt $(SRC_DIR)/...
