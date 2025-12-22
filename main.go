package main

import (
	"os"
	"envv/src/cmd/envv"
)

func main() {
	if err := envv.Execute(); err != nil {
		os.Exit(1)
	}
}
