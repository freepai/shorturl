# Makefile for building CoreDNS
GITCOMMIT:=$(shell git describe --dirty --always)
BINARY:=hummer
SYSTEM:=
CHECKS:=check
BUILDOPTS:=-v
GOPATH?=$(HOME)/go
MAKEPWD:=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
CGO_ENABLED:=0

.PHONY: all
all: hummer

.PHONY: hummer
hummer:
	CGO_ENABLED=$(CGO_ENABLED) $(SYSTEM) go build $(BUILDOPTS) -o $(BINARY)

.PHONY: test
test:
	go test ./...

.PHONY: cover
cover:
	go test -c -covermode=count -coverpkg ./...

.PHONY: run
run: hummer
	./hummer

.PHONY: clean
clean:
	go clean
	rm -f hummer