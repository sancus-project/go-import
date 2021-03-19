package app

import (
	"gopkg.in/gcfg.v1"

	"go.sancus.dev/middleware/goget"
)

// config.ini
type Config struct {
	HTTP struct {
		Address string
	}
	Package goget.Packages
}

func ConfigFromFile(fn string) (*Config, error) {
	var ini Config

	if err := gcfg.ReadFileInto(&ini, fn); err != nil {
		return nil, err
	}

	if err := ini.Package.SetDefaults(); err != nil {
		return nil, err
	}

	return &ini, nil
}
