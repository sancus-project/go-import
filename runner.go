package main

import (
	"app"
)

func main() {
	var fn = "config.ini"
	var l = app.Loggers.Get("go.sancus.io")

	s, err := app.NewServerFromFile(fn, l)
	if err != nil {
		l.Fatal("NewServerFromFile(%q): %s", fn, err)
	}

	l.Info("Listening %s", s.Addr)
	s.ListenAndServe()
}
