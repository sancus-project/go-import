package app

import (
	"fmt"
	"go.sancus.io/core/log"
	"go.sancus.io/web"
	"net/http"
	"strings"
)

// handler
type handler struct {
	packages map[string]string
	logger   *log.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.Host + r.URL.Path

	for k, v := range h.packages {
		s := strings.TrimPrefix(url, k)
		if s == "" || s[0] == '/' {
			if s == "" {
				h.logger.Info("%s -> %s", k, v)
			} else {
				h.logger.Info("%s (%s) -> %s", k, s, v)
			}

			fmt.Fprintf(w, "<!DOCTYPE html>\n<head>\n")
			fmt.Fprintf(w, "\t<meta name=\"go-import\" content=\"%s\">\n", v)
			fmt.Fprintf(w, "</head>\n<body />\n")
			return
		}
	}

	h.logger.Warn("%v: not recorgnized", url)
	http.NotFound(w, r)
	return
}

func NewHandler(packages map[string]*Package, l *log.Logger) http.Handler {
	h := handler{
		packages: make(map[string]string),
		logger:   l.SubLogger(":dispatcher"),
	}

	for k, v := range packages {
		if v.URL == "" {
			continue
		}
		if v.VCS == "" {
			v.VCS = "git"
		}

		s := "%s %s %s"
		h.packages[k] = fmt.Sprintf(s, k, v.VCS, v.URL)
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
