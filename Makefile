APPNAME   := $(shell basename $(CURDIR))
VERSION   := $(shell git describe --abbrev=0 --tags 2>/dev/null)
REVISION  := $(shell git rev-parse HEAD 2>/dev/null)

ifeq ($(VERSION),)
VERSION := dev
endif

ifeq ($(REVISION),)
REVISION := unknown
endif

LDFLAGS_APPNAME  := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(VERSION)"
LDFLAGS_REVISION := -X "main.Revision=$(REVISION)"
LDFLAGS          := -s -w -buildid= $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static
BUILDFLAGS       := -trimpath -ldflags '$(LDFLAGS)'

.PHONY: all
all: clean tools generate lint vet test build

.PHONY: tools
tools:
	aqua install --all

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: vet
vet:
	CGO_ENABLED=0 go vet ./...

.PHONY: test
test:
	CGO_ENABLED=0 go test ./...

.PHONY: build
build: bin/$(BINNAME)
bin/$(BINNAME): $(SRCS)
	CGO_ENABLED=0 go build $(BUILDFLAGS) -o $@

.PHONY: install
install: build
	CGO_ENABLED=0 go install $(BUILDFLAGS)

.PHONY: release
release:
ifneq ($(GITHUB_TOKEN),)
	goreleaser release --rm-dist
endif

.PHONY: snapshot
snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: clean
clean:
	rm -rf bin
	rm -rf dist
