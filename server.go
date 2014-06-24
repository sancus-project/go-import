package app

import (
	"fmt"
	"go.sancus.io/core/log"
	"go.sancus.io/web"
	"net/http"
	"regexp"
)

// handler
type handler struct {
	packages map[string]string
	logger   *log.Logger
}

var path_split = regexp.MustCompile("^/(([^/]+)(/.*)?)?$")

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := path_split.FindAllStringSubmatch(r.URL.Path, -1)[0][2]
	v := h.packages[p]

	if v == "" {
		h.logger.Warn("path: %s (not recorgnized)", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	h.logger.Info("path: %s (%s)", r.URL.Path, p)

	fmt.Fprintf(w, "<!DOCTYPE html>\n<head>\n")
	fmt.Fprintf(w, "\t<meta name=\"go-import\" content=\"%s\">\n", v)
	fmt.Fprintf(w, "</head>\n<body />\n")
}

func NewHandler(packages map[string]*Package, l *log.Logger) http.Handler {
	h := handler{
		packages: make(map[string]string),
		logger:   l.SubLogger(":dispatcher"),
	}

	for k, v := range packages {
		if v.Repo == "" {
			continue
		}
		if v.VCS == "" {
			v.VCS = "git"
		}

		s := "go.sancus.io/%s %s %s"
		h.packages[k] = fmt.Sprintf(s, k, v.VCS, v.Repo)
	}

	return &h
}

func NewServerFromFile(fn string, l *log.Logger) (*web.Server, error) {
	var err error
	if ini, err := ConfigFromFile(fn); err == nil {
		h := NewHandler(ini.Package, l)
		s := web.NewServer(ini.HTTP.Address, h)
		return s, nil
	}
	return nil, err
}
