package main

import (
	"log"

	"github.com/blushft/strana/cmd/cli/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
