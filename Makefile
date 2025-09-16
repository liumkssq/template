.PHONY: all fmt lint test build release deps install-tools check

GO=go
BINARY=myapp

all: build

install-tools:
	@./scripts/install-tools.sh

fmt:
	@gofumpt -w .
	@goimports -w .

lint:
	@golangci-lint run

test:
	@go test ./... -v

vet:
	@go vet ./...

build:
	@$(GO) build -o ./dist/$(BINARY) ./cmd/app

cross-build:
	@goreleaser build --rm-dist

changelog:
	@git-cliff -c .gitcliff.toml -o CHANGELOG.md

release: changelog
	@goreleaser release --rm-dist
