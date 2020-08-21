package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"task-tool-cli/cmd"
)

func main() {
	cmd, err := cmd.NewRootCmd()
	if err != nil {
		logrus.Debug(err)
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
