package cook

import (
	"io/ioutil"
	"runtime"
	"time"

	"github.com/mdayaram/cgofail/jello"
)

type Order struct {
	Jellos int
	Done   chan time.Duration
}

type Cook struct {
	order_up     chan *Order
	gel          jello.Jello
	jello_recipe string
}

func New(order_up chan *Order, jello jello.Jello, recipe string) *Cook {
	inbytes, err := ioutil.ReadFile(recipe)
	if err != nil {
		panic(err)
	}
	instr := string(inbytes)
	return &Cook{order_up: order_up, gel: jello, jello_recipe: instr}
}

func (w *Cook) StartCooking(use_private_kitchen bool) {
	go func() {
		if use_private_kitchen {
			runtime.LockOSThread()
		}

		var start time.Time
		var dur time.Duration
		for {
			order := <-w.order_up
			start = time.Now()
			for i := 0; i < order.Jellos; i++ {
				jresult := w.gel.Jiggle(w.jello_recipe, w.jello_recipe)
				if jresult != w.jello_recipe+w.jello_recipe {
					panic("Customer found a bug in their Jello.")
				}
			}
			dur = time.Now().Sub(start)
			order.Done <- dur
		}
	}()
}
