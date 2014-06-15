.PHONY: run

WS := $(CURDIR)
GO := go

T := $(WS)/runner.go

GOPATH := $(WS)$(shell echo "$${GOPATH:+:$$GOPATH}")
GOBIN  := $(WS)/bin
TMPDIR := $(WS)/tmp

export GOPATH GOBIN TMPDIR

run: prepare
	$(GO) run $(T)

all:
	$(GO) install -x $(T)

prepare: $(TMPDIR)

$(TMPDIR):
	mkdir -p $(TMPDIR)
