package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

var (
	flagParser = flags.NewNamedParser("srclib-man", flags.Default)
	cwd        = getCWD()
)

func init() {
	flagParser.LongDescription = "srclib-man extracts defs of POSIX commands for srclib-bash."
}

func getCWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return cwd
}

func main() {
	log.SetFlags(0)
	if _, err := flagParser.Parse(); err != nil {
		os.Exit(1)
	}
}
