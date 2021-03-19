.PHONY: build fmt tidy clean

GO ?= go
GOFMT ?= gofmt -w -l -s
GOGET ?= $(GO) get -v

build:
	$(GOGET) ./...

tidy:
	$(GO) mod tidy

fmt: tidy
	find -name '*.go' | xargs -r $(GOFMT)

clean: tidy
	$(GO) clean -x -r -modcache
