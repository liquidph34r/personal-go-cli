package liquid

import (
	"flag"
	"fmt"
	"os"
)

//https://developers.google.com/youtube/v3/guides/implementation/search

func startController() {
	// Showing useful information when the user enters the --help option
	print("Liquid v1.0\n\n")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	cmd := os.Args[1]

	switch cmd {
	case "searchvideo":
		GetSearchVideoFlags(os.Args[2:])
	default:
		flag.Usage()
	}
}
