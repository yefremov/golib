#
# golib
#

GO ?= go
GOVET := ${GO} vet
GOFMT := gofmt -l
GOTESTFLAGS ?= --cover --race -v
GOTEST := ${GO} test ${GOTESTFLAGS}
GOCLEAN := ${GO} clean

vet:
	@${GOVET}
.PHONY: vet

fmt:
	@${GOFMT} -w .
.PHONY: fmt

test:
	@${GOTEST} ./...
.PHONY: test

clean:
	@${GOCLEAN} .
.PHONY: clean

all: vet fmt test
.PHONY: all

.DEFAULT_GOAL := all
