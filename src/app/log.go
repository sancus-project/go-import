package app

import (
	"go.sancus.io/core/log"
)

var Loggers = log.NewGroup(log.INFO, &log.StderrBackend)
