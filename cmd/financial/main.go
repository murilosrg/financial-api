package main

import (
	"flag"
)

func main() {
	load()
}

func load() {
	var shouldInit = flag.Bool("init", false, "initialize all")
	flag.Parse()

	if *shouldInit {
		initAll()
	}
}
