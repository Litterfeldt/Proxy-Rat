package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultPort = "8080"
)

var port = flag.String("port", defaultPort, "The port number")

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nCommand line arguments:\n\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(2)
	}
	flag.Parse()
}
