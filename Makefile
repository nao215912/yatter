BINARY := yatter-backend-go
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

PATH := $(PATH):${MAKEFILE_DIR}bin
SHELL := env PATH="$(PATH)" /bin/bash
# for go
export CGO_ENABLED = 0
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_URL=local-git://

LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

build: build-linux

build-default:
	go build ${LDFLAGS} -o build/${BINARY}

build-linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o build/${BINARY}-linux-${GOARCH} .

prepare: mod

mod:
	go mod download

test:
	go test -v -count=1 $(shell go list ${MAKEFILE_DIR}/...)

lint:
	if ! [ -x $(GOPATH)/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.38.0 ; \
	fi
	golangci-lint run --concurrency 2

vet:
	go vet ./...

clean:
	git clean -f -X app bin build

up:
	docker-compose up

stop:
	docker-compose stop


del:
	docker system prune -af
	rm -rf .data

web:
	docker-compose exec web bash

db:
	docker-compose exec db mysql -u root


.PHONY:	test clean up stop del
