COMMIT = $(shell git rev-list --count --all)
BRANCH = $(shell git symbolic-ref -q HEAD | cut -b 12-)
VERSION = $(BRANCH)-$(shell git rev-parse --short HEAD)

LDFLAGS += -X 'gocms/handler.build=$(COMMIT)'

.PHONY: all

all: vendor node_modules build

vendor:
	@GO111MODULE=on go mod vendor

node_modules:
	@npm install

build:
	@echo "building: ${VERSION}"
	@npm run build
	@go build -v -ldflags "$(LDFLAGS)"