.PHONY: default

GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
BINARY=barkme

FLAGS="-s -w"

default:
	@echo "build target is required"
	@exit 1

build:
	$(GOBUILD) -v -trimpath -ldflags $(FLAGS) -o $(BINARY)
build_linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -trimpath -ldflags $(FLAGS) -o $(BINARY)
build_mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -v -trimpath -ldflags $(FLAGS) -o $(BINARY)
build_win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -v -trimpath -ldflags $(FLAGS) -o $(BINARY).exe

install:
	$(GOINSTALL) -v -trimpath -ldflags $(FLAGS) .
preview: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
	rm -f $(BINARY).exe
