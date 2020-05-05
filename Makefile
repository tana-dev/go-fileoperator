# Meta Information
NAME := go-fileoperator
BINARY_LINUX := $(NAME)_linux
BINARY_WINDOWS := $(NAME).exe
BINARY_MACOS := $(NAME)_macos
GOARCH := amd64

# Setup
.PHONY: setup
setup:
	go get golang.org/x/tools/cmd/...
	go get golang.org/x/lint/golint

# Test
## lint
.PHONY: lint
lint:
	go vet ./...
	golint -set_exit_status ./...

## test
.PHONY: test
test:
	goimports -w ./
	go test ./...

# Clean
.PHONY: clean
clean:
	-rm -rdf ./statik

# Build
## statik
.PHONY: statik
statik:
	-rm -rdf ./statik
	statik -src static

## build-linux
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -o ./cmd/$(BINARY_LINUX) -v ./cmd/$(NAME)/main.go

## build-windows
.PHONY: build-windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) go build -o ./cmd/$(BINARY_WINDOWS) -v ./cmd/$(NAME)/main.go

## build-macos
.PHONY: build-macos
build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=$(GOARCH) go build -o ./cmd/$(BINARY_MACOS) -v ./cmd/$(NAME)/main.go

## all build
.PHONY: build
build: statik build-linux build-windows build-macos clean
