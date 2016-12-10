.PHONY: default test lint buld build-all package clean

BIN=snpt
BUILD_DIR=build

default: build

test:
	go test -v ./

lint:
	gometalinter --disable=errcheck ./

build:
	go build -o $(BUILD_DIR)/$(BIN) ./

build-all:
	gox -output "$(BUILD_DIR)/$(BIN)-$(TRAVIS_TAG)-{{.OS}}-{{.Arch}}/$(BIN)" ./

package:
	find ./$(BUILD_DIR) -mindepth 1 -maxdepth 1 -type d -execdir tar -czvf {}.tar.gz {} \;

clean:
	rm -rf ./$(BUILD_DIR)

install:
	go get -t -v ./...

install-env-deps:
	go get -u github.com/mitchellh/gox
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
