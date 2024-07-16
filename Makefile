ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
export BUILD_PATH ?= $(ROOT_DIR)/_build

VERSION := $(shell git describe --tags --always)
ARCH := $(shell go env GOARCH)
OS := $(shell go env GOOS)
export PACKAGE_NAME := webservice-$(VERSION)-$(OS)-$(ARCH)

.PHONY: build
build: clean build_prepare
	go build -o webservice cmd/main.go
	@mv ./webservice $(BUILD_PATH)/$(PACKAGE_NAME)/bin
	@echo "Build successfully"

.PHONY: build_prepare
build_prepare:
	@mkdir -p $(BUILD_PATH)/$(PACKAGE_NAME)/bin

.PHONY: clean
clean:
	@rm -rf _build
