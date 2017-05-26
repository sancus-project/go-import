package app

import (
	"gopkg.in/gcfg.v1"
)

// config.ini
type Package struct {
	VCS  string
	URL string
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
