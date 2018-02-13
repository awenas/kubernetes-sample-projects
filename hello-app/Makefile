IMAGE = "hello-app"
VERSION = $(shell git describe --always)
USERNAME = "owainlewis"

all: build build-docker

.PHONY: build
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/hello-app main.go

.PHONY: build-docker
build-docker:
	@docker build -t $(IMAGE):$(VERSION) .

.PHONY: push
push:
	@docker push $(USERNAME)/$(IMAGE):$(VERSION)
