CI_COMMIT_REF_NAME ?= $(shell git symbolic-ref -q HEAD | cut -b 12-)
CI_COMMIT_SHORT_SHA ?= $(shell git rev-parse --short HEAD)
CI_REGISTRY ?= registry.loc:5000

LDFLAGS += -X 'gocms/handler.build=$(CI_COMMIT_SHORT_SHA)' -w -s

.PHONY: all

all: build

node_modules:
	@npm install

front: node_modules
	@npm run build

build:
	@echo "building: $(CI_COMMIT_REF_NAME)"
	@go build -v -ldflags "$(LDFLAGS)"