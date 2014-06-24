package app

import (
	"code.google.com/p/gcfg"
)

// config.ini
type Package struct {
	VCS  string
	Repo string
}

type Config struct {
	HTTP struct {
		Address string
	}
	Package map[string]*Package
}

func ConfigFromFile(fn string) (*Config, error) {
	var ini Config

	if err := gcfg.ReadFileInto(&ini, fn); err != nil {
		return nil, err
	}

	return &ini, nil
}
