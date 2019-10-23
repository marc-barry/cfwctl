DOCKER_ORG := middlenamesfirst
NAME := cfwctl
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
BASE_IMAGE_URL := ${DOCKER_ORG}/$(NAME)
IMAGE_URL := $(BASE_IMAGE_URL):$(GIT_COMMIT)

.PHONY: build
build: go-modules-tidy
	go build -ldflags "-X github.com/marc-barry/cfwctl/version.GitCommit=$(GIT_COMMIT)" -o bin/cfwctl cmd/cfwctl/main.go

.PHONY: docker-build
docker-build:
	docker build --pull -t ${IMAGE_URL} .

.PHONY: docker-push
docker-push: docker-build
	docker push "${IMAGE_URL}"

.PHONY: go-modules-tidy
test: go-modules-tidy
	go test --count=1 -cover ./...

.PHONY: go-modules-tidy
go-modules-tidy:
	go mod tidy
