package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func proc(ch <-chan string, watchName, command string) {
	for changed := range ch {
		run(path.Join(watchName, changed), command)
	}
}

func run(changedName, command string) {
	command = os.Expand(command, func(s string) string {
		if s == "roc_name" {
			return changedName
		}
		return os.Getenv(s)
	})
	logPrintf(command)
	cmd := exec.Command("zsh", "-c", command)
	if err := cmd.Run(); err != nil {
		log.Printf("command failed: %v", err)
		return
	}
}
