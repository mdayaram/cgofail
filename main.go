package main

import (
	"flag"
	"github.com/mdayaram/cgofail/worker"
)

func main() {
	workers := flag.Int("w", 1, "Number of workers to allocate.")
	flag.Parse()

	worker.What()
	println("Number of workers: ", *workers)
}
