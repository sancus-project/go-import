.PHONY: run all

WS := $(CURDIR)
GO := go

T := $(WS)/runner.go

DEPS := $(patsubst %,src/%/.git, code.google.com/p/gcfg \
	go.sancus.io/core \
	go.sancus.io/web)

GOPATH := $(WS)$(shell echo "$${GOPATH:+:$$GOPATH}")
GOBIN  := $(WS)/bin
TMPDIR := $(WS)/tmp

export GOPATH GOBIN TMPDIR

run: prepare
	$(GO) run $(T)

all:
	$(GO) install -x $(T)

prepare: $(TMPDIR) $(DEPS)

$(DEPS): D=$(patsubst src/%/.git,%,$@)
$(DEPS):
	$(GO) get $(D)

$(TMPDIR):
	mkdir -p $(TMPDIR)
