DIST := dist
APPNAME := halsecur

GOLANGCILINT_VERSION := v2.7.2
GOSEC_VERSION := v2.22.11
VULNCHECK_VERSION := latest

LDFLAGS := -X bisecur/version.Version=?version? -X bisecur/version.BuildDate=?date?

ifeq ($(OS),Windows_NT)
	SHELL := cmd.exe
	SHELLFLAGS := /C

	EXE := .exe
	DEVNULL := NUL

	# Use cmd built-ins / Windows tools
	MKDIR_P = if not exist "$(DIST)" mkdir "$(DIST)"
	RM_RF   = if exist "$(DIST)" rmdir /S /Q "$(DIST)"
	WHERE   = where
else
	EXE :=
	DEVNULL := /dev/null

	MKDIR_P = mkdir -p "$(DIST)"
	RM_RF   = rm -rf "$(DIST)"
	WHERE   = which
endif

OUT := $(DIST)/$(APPNAME)$(EXE)

# --- Targets ---
.PHONY: all env clean lint-env lint lint-fix test test-short build build-linux build-docker

all: clean build build-linux

env:
	@$(MKDIR_P)

clean:
	@$(RM_RF)

lint-env:
	@$(WHERE) gosec >$(DEVNULL) 2>&1 && gosec --version | grep -qs "$(GOSEC_VERSION)" || go install github.com/securego/gosec/v2/cmd/gosec@$(GOSEC_VERSION)
	@$(WHERE) golangci-lint >$(DEVNULL) 2>&1 && golangci-lint --version | grep -qs "$(GOLANGCILINT_VERSION)" || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCILINT_VERSION)
	@$(WHERE) govulncheck >$(DEVNULL) 2>&1 || go install golang.org/x/vuln/cmd/govulncheck@$(VULNCHECK_VERSION)

lint: lint-env
	golangci-lint --timeout 10m -v run ./...
	gosec ./...
	govulncheck ./...

lint-fix: lint-env
	golangci-lint run -v --fix ./...

test: test-short
	go test $(VENDOR) ./...

test-short:
	go test $(VENDOR) -race -short ./...

build: env
ifeq ($(OS),Windows_NT)
	@set CGO_ENABLED=0&& go build -ldflags "$(LDFLAGS)" -v -o "$(OUT)" .
else
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -v -o "$(OUT)" .
endif

build-linux: env
ifeq ($(OS),Windows_NT)

	@set CGO_ENABLED=0&& @set GOARCH=amd64&& @set GOOS=linux&& go build -ldflags "$(LDFLAGS)" -v -o "$(DIST)/$(APPNAME)" .
else
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -v -o "$(DIST)/$(APPNAME)" .
endif

build-docker: env build
	docker build --build-arg TARGETPLATFORM='./' --build-arg VERSION=$(shell git describe --tags --always) -t bisecur/halsecur:latest -f Dockerfile "$(DIST)"
