package main

import (
	"os"

	"github.com/liepertt/serialbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
