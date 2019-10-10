GIT_COMMIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null)

.PHONY: build
build: go-modules-tidy
	go build -ldflags "-X github.com/marc-barry/cfwctl/version.GitCommit=$(GIT_COMMIT_SHA)" -o bin/cfwctl cmd/cfwctl/main.go

.PHONY: go-modules-tidy
go-modules-tidy:
	go mod tidy
