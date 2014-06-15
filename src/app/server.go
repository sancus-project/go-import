package app

import (
	"fmt"
	"go.sancus.io/core/log"
	"go.sancus.io/web"
	"net/http"
)

// handler
type handler struct {
	projects map[string]string
	logger   *log.Logger
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("path: %s", r.URL.Path)
	http.NotFound(w, r)
}

func NewProjectHandler(projects map[string]*GoImport, l *log.Logger) http.Handler {
	h := handler{
		projects: make(map[string]string),
		logger:   l.SubLogger("dispatcher"),
	}

	for k, v := range projects {
		if v.Repo == "" {
			continue
		}
		if v.VCS == "" {
			v.VCS = "git"
		}

		s := "go.sancus.io/%s %s %s"
		h.projects[k] = fmt.Sprintf(s, k, v.VCS, v.Repo)
	}

	return &h
}

func NewServerFromFile(fn string, l *log.Logger) (*web.Server, error) {
	if ini, err := ConfigFromFile(fn); err == nil {
		h := NewProjectHandler(ini.Project, l)
		s := web.NewServer(ini.HTTP.Address, h)
		return s, nil
	} else {
		return nil, err
	}
}
