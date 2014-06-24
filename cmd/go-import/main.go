package main

import (
	"go.sancus.io/core/log"
	"go.sancus.io/go-import"
)

func main() {
	var fn = "config.ini"
	var l = log.GetLogger("go-import")

	s, err := app.NewServerFromFile(fn, l)
	if err != nil {
		l.Fatal("NewServerFromFile(%q): %s", fn, err)
	}

	l.Info("Listening %s", s.Addr)
	s.ListenAndServe()
}
