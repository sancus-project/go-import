.PHONY: run all

WS := $(CURDIR)
GO := go
T  := $(WS)/runner.go

DEPS := code.google.com/p/gcfg \
	go.sancus.io/core \
	go.sancus.io/web

SDEPS := $(patsubst %,src/%/.git, $(DEPS))

GOPATH := $(WS)$(shell echo "$${GOPATH:+:$$GOPATH}")
GOBIN  := $(WS)/bin
TMPDIR := $(WS)/tmp

export GOPATH GOBIN TMPDIR

run: prepare
	$(GO) run $(T)

all: $(SDEPS)
	$(GO) install -x $(T)

prepare: $(TMPDIR) $(DEPS)

$(SDEPS): D=$(patsubst src/%/.git,%,$@)
$(SDEPS):
	$(GO) get $(D)

$(TMPDIR):
	mkdir -p $(TMPDIR)
