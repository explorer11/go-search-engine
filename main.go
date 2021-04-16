package main

import (
	"flag"
	"fmt"
	"searchengine/commandline"
	"searchengine/http"
)

func main() {

	var mode string
	var directory string
	flag.StringVar(&mode, "m", "", "mode http or command line")
	flag.StringVar(&directory, "d", "", "specify directory to use")
	flag.Parse()
	fmt.Printf("mode = %s\n", mode)
	fmt.Printf("directory = %s\n", directory)

	if mode != "command" {
		http.Run(directory)
	} else {
		commandline.Run(directory)
	}
}
