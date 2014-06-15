package app

import (
	"code.google.com/p/gcfg"
)

// config.ini
type GoImport struct {
	VCS  string
	Repo string
}

type Config struct {
	HTTP struct {
		Address string
	}
	Project map[string]*GoImport
}

func ConfigFromFile(fn string) (*Config, error) {
	var ini Config

	if err := gcfg.ReadFileInto(&ini, fn); err != nil {
		return nil, err
	}

	return &ini, nil
}
