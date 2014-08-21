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
	packages map[string]*Package
	logger   *log.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.Host + r.URL.Path

	for k, v := range h.packages {
		s := strings.TrimPrefix(url, k)
		if s == "" || s[0] == '/' {
			if s == "" {
				h.logger.Info("%s -> %s+%s", k, v.VCS, v.URL)
			} else {
				h.logger.Info("%s (%s) -> %s+%s", k, s, v.VCS, v.URL)
			}

			s = "https://godoc.org/" + url

			fmt.Fprintf(w, "<!DOCTYPE html>\n<head>\n")
			fmt.Fprintf(w, "\t<meta name=\"go-import\" content=\"%v %v %v\">\n", k, v.VCS, v.URL)
			fmt.Fprintf(w, "\t<meta http-equiv=\"refresh\" content=\"5; url=%v\" />\n", s)

			fmt.Fprintf(w, "</head>\n<body>\n")
			fmt.Fprintf(w, "<pre>git clone <a href=\"%v\">%v</a>\n</pre>", v.URL, v.URL)
			fmt.Fprintf(w, "<pre>go get <a href=\"%v\">%v</a></pre>\n", s, url)
			fmt.Fprintf(w, "<pre>import \"<a href=\"%v\">%v</a>\"</pre>\n", s, url)
			fmt.Fprintf(w, "</body>\n")
			return
		}
	}

	h.logger.Warn("%v: not recognised", url)
	http.NotFound(w, r)
	return
}

func NewHandler(packages map[string]*Package, l *log.Logger) http.Handler {
	h := handler{
		packages: make(map[string]*Package),
		logger:   l.SubLogger(":dispatcher"),
	}

	for k, v := range packages {
		if v.URL == "" {
			continue
		}
		if v.VCS == "" {
			v.VCS = "git"
		}

		h.packages[k] = v
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
