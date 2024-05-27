package main

import (
	"animatic-v2/internal/argparse"
	"animatic-v2/internal/debug"
	"animatic-v2/internal/orchester"
	"flag"
)

func main() {
	argv, err := argparse.Parse(flag.CommandLine)
	if err != nil {
		debug.LogError("parsing arguments", err)
	}

	orchester.Start(argv)
}
