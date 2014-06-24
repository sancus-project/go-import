.PHONY: run all clean get fmt prepare

WS := $(CURDIR)
GO := go

PKG := go.sancus.io/go-import/cmd/go-import

GOPATH := $(WS)$(shell echo "$${GOPATH:+:$$GOPATH}")
GOBIN  := $(WS)/bin
TMPDIR := $(WS)/tmp

export GOPATH GOBIN TMPDIR

PKGDIR := $(WS)/pkg
DIRS   := $(PKGDIR) $(TMPDIR) $(GOBIN)

all: $(DIRS)
	$(GO) install -v $(PKG)

get: $(DIRS)
	$(GO) get -u $(PKG)

run: all
	$(GOBIN)/go-import

fmt:
	find $(WS) -name 'gofmt.sh' -exec $(SHELL) '{}' \;

$(DIRS):
	mkdir $@

clean:
	rm -vrf $(DIRS)
