#! /usr/bin/make
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "gen" invokes the "goa" tool to generate the examples source code
# - "build" compiles the example microservices and client CLIs
# - "clean" deletes the output of "build"
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
GO_FILES=$(shell find . -type f -name '*.go')
GOA:=$(shell goa version 2> /dev/null)

.PHONY: all gen-goa

all: check-goa gen-goa tidy format

check-goa:
ifdef GOA
	@echo $(GOA)
else
	go install goa.design/goa/v3/...@v3
	@echo $(GOA)
endif

gen-goa:
	@echo GENERATING CODE...
	@goa version
	@rm -rf "./gen"
	@goa gen goa/design -o "./"

tidy:
	@echo TIDYING CODE...
	@go mod tidy -compat=1.17

format:
	@echo FORMAT PROJECT CODE...
	@go fmt ./...

