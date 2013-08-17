package main

import (
	"flag"
	"sync"
	"time"

	"github.com/mdayaram/cgofail/cook"
	"github.com/mdayaram/cgofail/jello"
	"github.com/mdayaram/simstat"
)

var wait4me sync.WaitGroup
var datas = simstat.NewDataSet()

func main() {
	recipe := flag.String("r", "default_recipe.txt", "Jello recipe to use for making the jello.")
	jellos := flag.Int("j", 1, "Number of jellos a cook has to make for a single order.")
	cooks := flag.Int("c", 1, "Number of cookes in the kitcken cooking.")
	waiters := flag.Int("w", 100, "Number of waiters taking orders.")
	orders := flag.Int("o", 100, "Number of orders each waiter takes.")
	cgo := flag.Bool("cgo", false, "Use cgo for the jiggling instead of go.")
	lock := flag.Bool("lock", false, "Lock each cook to an OS thread.")
	flag.Parse()

	order_in := make(chan int, *cooks)
	order_out := make(chan time.Duration, *cooks)

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

	// Cooks! Prepare to start cooking!
	for i := 0; i < *cooks; i++ {
		var gel jello.Jello
		if *cgo {
			gel = jello.NewCgo()
		} else {
			gel = jello.NewGor()
		}
		c := cook.New(order_in, order_out, gel, *recipe)
		c.StartCooking(*lock)
	}

	// Waiters! Start taking orders!
	for i := 0; i < *waiters; i++ {
		wait4me.Add(1)
		go func() {
			for i := 0; i < *orders; i++ {
				order_in <- *jellos
				dur := <-order_out
				datas.Add(int64(dur))
			}
			wait4me.Done()
		}()
	}
	wait4me.Wait()

	println("Summary (units in nano seconds):\n")
	println(datas.String())
}
