package main

import (
	"go.sancus.dev/go-import"
	"go.sancus.dev/sancus/attic/log"
)

func main() {
	var fn = "config.ini"
	var l = log.GetLogger("go-import")

	s, err := app.NewServerFromFile(fn, l)
	if s == nil {
		l.Fatal("NewServerFromFile(%q): %v", fn, err)
	}

	l.Info("Listening %s", s.Addr)
	s.ListenAndServe()
}
