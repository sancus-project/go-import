package app

import (
	"net/http"
	"strings"

	"go.sancus.dev/middleware/goget"
	"go.sancus.dev/sancus/attic/log"
	"go.sancus.dev/sancus/web"
)

// handler
type handler struct {
	renderer goget.Renderer
	packages goget.Packages
	logger   *log.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	url := r.Host + r.URL.Path

	if v := h.packages.Get(r); v != nil {
		k := v.Canonical
		s := strings.TrimPrefix(url, k)

		if s == "" {
			h.logger.Info("%s -> %s+%s", k, v.VCS, v.Repository)
		} else {
			h.logger.Info("%s (%s) -> %s+%s", k, s, v.VCS, v.Repository)
		}

		if err := h.renderer(w, r, v); err != nil {
			h.logger.Fatal("%v: %s", url, err)
		}
	} else {

		h.logger.Warn("%v: not recognised", url)
		http.NotFound(w, r)
	}
}

func NewHandler(packages goget.Packages, l *log.Logger) http.Handler {
	h := handler{
		packages: packages,
		renderer: goget.DefaultRenderer(),
		logger:   l.SubLogger(":dispatcher"),
	}

	return &h
}

func NewServerFromFile(fn string, l *log.Logger) (*web.Server, error) {
	if ini, err := ConfigFromFile(fn); err == nil {
		h := NewHandler(ini.Package, l)
		s := web.NewServer(ini.HTTP.Address, h)
		return s, nil
	} else {
		return nil, err
	}
}
