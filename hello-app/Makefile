IMAGE = "hello-app"
VERSION = $(shell git describe --always)
USERNAME = "owainlewis"
RELEASE=0.1.0

all: build build-docker

.PHONY: build
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(VERSION)" -o ./bin/hello-app main.go

.PHONY: build-docker
build-docker: build
	@docker build -t $(IMAGE):$(VERSION) .

.PHONY: run-docker
run-docker:	build-docker
	docker run -p "9000:9000" -d $(IMAGE):$(VERSION)

.PHONY: push
push: build-docker
	@docker tag $(IMAGE):$(VERSION) $(USERNAME)/$(IMAGE):$(VERSION)
	@docker push $(USERNAME)/$(IMAGE):$(VERSION)

.PHONY: release
release: build-docker
	@docker tag $(IMAGE):$(VERSION) $(USERNAME)/$(IMAGE):$(RELEASE)
	@docker push $(USERNAME)/$(IMAGE):$(RELEASE)

	@docker tag $(IMAGE):$(VERSION) $(USERNAME)/$(IMAGE):latest
	@docker push $(USERNAME)/$(IMAGE):latest
