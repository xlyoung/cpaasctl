package main

import (
	"gitlab.hycyg.com/paas-tools/cpaasctl/cmd"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/logger"
)

var version string

func main() {
	logger.Logger.Infof("version: %s", version)
	cmd.Execute()
}
