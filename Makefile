COMMIT = $(shell git rev-list --count --all)
BRANCH = $(shell git symbolic-ref -q HEAD | cut -b 12-)
VERSION = $(BRANCH)-$(shell git rev-parse --short HEAD)

LDFLAGS += -X 'gocms/handler.build=$(COMMIT)'

.PHONY: all

all: build

node_modules:
	@npm install

static/js/admin.js:
	@npm run build

bindata.go: static/js/admin.js
	@go run github.com/go-bindata/go-bindata/go-bindata

build: bindata.go
	@echo "building: ${VERSION}"
	@go build -v -ldflags "$(LDFLAGS)"