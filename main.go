package main

import (
	"flag"
	"github.com/mdayaram/cgofail/jello"
	"github.com/mdayaram/cgofail/worker"
	"github.com/mdayaram/simstat"
	"sync"
	"time"
)

var wait4me sync.WaitGroup
var datas = simstat.NewDataSet()

func Provide(wch chan int, rch chan time.Duration, trials int) {
	for i := 0; i < trials; i++ {
		wch <- i
		dur := <-rch
		datas.Add(int64(dur))
	}
	wait4me.Done()
}

func main() {
	jiggles := flag.Int("j", 1, "Number of times to jiggle the jello (# of times work is executed).") // Jiggles
	workers := flag.Int("c", 1, "Number of workers to allocate.")                                     // Cooks
	concurrency := flag.Int("w", 100, "Number of work providers (concurrency).")                      // Waiters
	trials := flag.Int("d", 100, "Number of pieces of work each provider provides (trials).")         // Diners
	cgo := flag.Bool("cgo", false, "Use cgo for the jiggling instead of go.")
	lock := flag.Bool("lock", false, "Lock each worker to an OS thread.")
	flag.Parse()

	sendWork := make(chan int, *workers)
	recvResults := make(chan time.Duration, *workers)

	if *cgo {
		println("Using cgo jello...")
	} else {
		println("Using normal jello...")
	}
	println("Hiring", *workers, "[c]ooks...")
	if *lock {
		println(" + With their dedicated kitchens...")
	}
	println("Hiring", *concurrency, "[w]aiters...")
	println("Each w", *trials, "[d]iners...")
	println("Giving each jello", *jiggles, " [j]iggles...")

	for i := 0; i < *workers; i++ {
		var gel jello.Jello
		if *cgo {
			gel = jello.NewCgo()
		} else {
			gel = jello.NewGor()
		}
		w := worker.New(sendWork, recvResults, gel, *jiggles)
		w.WorkIt(*lock)
	}

	for i := 0; i < *concurrency; i++ {
		wait4me.Add(1)
		go Provide(sendWork, recvResults, *trials)
	}
	wait4me.Wait()

	println("Summary (units in nano seconds):\n")
	println(datas.String())
}
