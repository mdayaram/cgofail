package main

import (
	"flag"
	"sync"
	"time"

	"github.com/mdayaram/cgofail/jello"
	"github.com/mdayaram/cgofail/worker"
	"github.com/mdayaram/simstat"
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
	recipe := flag.String("r", "default_recipe.txt", "Jello recipe to use for making the jello.")
	jellos := flag.Int("j", 1, "Number of jellos a cook has to make for a single order.")
	cooks := flag.Int("c", 1, "Number of cookes in the kitcken cooking.")
	waiters := flag.Int("w", 100, "Number of waiters taking orders.")
	orders := flag.Int("o", 100, "Number of orders each waiter takes.")
	cgo := flag.Bool("cgo", false, "Use cgo for the jiggling instead of go.")
	lock := flag.Bool("lock", false, "Lock each cook to an OS thread.")
	flag.Parse()

	sendWork := make(chan int, *cooks)
	recvResults := make(chan time.Duration, *cooks)

	if *cgo {
		println("Using [cgo] jello...")
	} else {
		println("Using normal jello...")
	}
	println("Made from the", *recipe, "[r]ecipe...")
	println("Hiring", *cooks, "[c]ooks...")
	if *lock {
		println(" + Each [lock]ed to their dedicated kitchens...")
	}
	println("Hiring", *waiters, "[w]aiters...")
	println("Each taking", *orders, "[o]rders...")
	println("Each order requiring", *jellos, "[j]ellos...")

	for i := 0; i < *cooks; i++ {
		var gel jello.Jello
		if *cgo {
			gel = jello.NewCgo()
		} else {
			gel = jello.NewGor()
		}
		w := worker.New(sendWork, recvResults, gel, *jellos, *recipe)
		w.WorkIt(*lock)
	}

	for i := 0; i < *waiters; i++ {
		wait4me.Add(1)
		go Provide(sendWork, recvResults, *orders)
	}
	wait4me.Wait()

	println("Summary (units in nano seconds):\n")
	println(datas.String())
}
