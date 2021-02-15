.PHONY: default

GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
BINARY=barkme

default:
	@echo "build target is required"
	@exit 100
build:
	$(GOBUILD) -v -ldflags "-s -w" -o $(BINARY) .
install:
	$(GOINSTALL) -v -ldflags "-s -w" .
preview: build
	./$(BINARY)
