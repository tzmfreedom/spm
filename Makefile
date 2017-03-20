NAME := spm
SRCS := $(shell find . -type d -name vendor -prune -o -type f -name "*.go" -print)
VERSION := 0.1.2
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\"" 

.DEFAULT_GOAL := bin/$(NAME) 

.PHONY: test
test: glide
	@go test -cover -v `glide novendor`

.PHONY: install
install: build
	@go install

.PHONY: uninstall
uninstall:

.PHONY: clean
clean:
	@rm -rf bin/*
	@rm -rf vendor/*
	@rm -rf dist/*

.PHONY: dist-clean
dist-clean: clean
	@rm -f $(NAME).tar.gz

.PHONY: build
build: 
	-@goimports -w $(SRCS)
	@gofmt -w $(SRCS)
	@go build $(LDFLAGS)

.PHONY: cross-build
cross-build: deps
	-@goimports -w $(SRCS)
	@gofmt -w $(SRCS)
	@for os in darwin linux windows; do \
	    for arch in amd64 386; do \
	        GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo \
	        -installsuffix netgo $(LDFLAGS) -o dist/$(NAME)-$$os-$$arch; \
	    done; \
	done

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: deps
deps: glide
	glide install

.PHONY: bin/$(NAME) 
bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: dist
dist:
	@tar czfh $(NAME).tar.gz $(shell git ls-files)
