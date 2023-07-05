package main

import (
	"os"

	"docker/container"

	log "github.com/sirupsen/logrus"
)

// Run command实际执行的命令
func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
