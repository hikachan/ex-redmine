SRCS = $(shell git ls-files '*.go')

## Setup
setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get github.com/Songmu/make2help/cmd/make2help
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/exp/utf8string

## Run tests
test: deps
	go test $$(glide novendor)

## glide setup
deps: setup
	glide install

## glide update
update: setup
	glide update

## Lint
lint:
	@ go get -v github.com/golang/lint/golint
	$(foreach file,$(SRCS),golint $(file) || exit;)

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup test deps update lint help
