VERSION ?= $$(git describe --tags)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/icon

release: build
	tar czvf icon.tar.gz bin/
