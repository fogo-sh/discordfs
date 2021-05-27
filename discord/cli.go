package discord

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s ROOTDIR\n", os.Args[0])
	flag.PrintDefaults()
}

// Cli wraps the discord file system as a command line interface.
func Cli() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	token := os.Getenv("DISCORDFS_TOKEN")
	rootdir := flag.Arg(0)

	Run(token, rootdir)
}
