package main

import (
	"flag"
	"github.com/mdayaram/cgofail/worker"
	"sync"
	"time"
)

var wait4me sync.WaitGroup

func Provide(wch chan int, rch chan time.Duration, trials int) {
	for i := 0; i < trials; i++ {
		wch <- i
		dur := <-rch
		println(dur / 1000)
	}
	wait4me.Done()
}

func main() {
	workers := flag.Int("w", 1, "Number of workers to allocate.")
	concurrency := flag.Int("c", 10000, "Number of work providers (concurrency).")
	trials := flag.Int("t", 10000, "Number of pieces of work each provider provides (trials).")
	flag.Parse()

	sendWork := make(chan int, *workers)
	recvResults := make(chan time.Duration, *workers)

	for i := 0; i < *workers; i++ {
		w := worker.New(sendWork, recvResults)
		w.WorkIt()
	}

	for i := 0; i < *concurrency; i++ {
		wait4me.Add(1)
		go Provide(sendWork, recvResults, *trials)
	}
	wait4me.Wait()
}
