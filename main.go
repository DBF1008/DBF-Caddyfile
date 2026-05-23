package main

import (
	"os"

	"github.com/hashicorp/boundary/internal/cmd"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"dev", "-container-image", "docker.m.daocloud.io/library/postgres:15"}
	}
	os.Exit(cmd.Run(args))
}