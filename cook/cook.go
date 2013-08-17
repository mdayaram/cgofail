package cook

import (
	"io/ioutil"
	"runtime"
	"time"

	"github.com/mdayaram/cgofail/jello"
)

type Cook struct {
	recv         chan int
	send         chan time.Duration
	gel          jello.Jello
	jello_recipe string
}

func New(recv chan int, send chan time.Duration, jello jello.Jello, recipe string) *Cook {
	inbytes, err := ioutil.ReadFile(recipe)
	if err != nil {
		panic(err)
	}
	instr := string(inbytes)
	return &Cook{recv: recv, send: send, gel: jello, jello_recipe: instr}
}

func (w *Cook) StartCooking(use_private_kitchen bool) {
	go func() {
		if use_private_kitchen {
			runtime.LockOSThread()
		}

		var start time.Time
		var dur time.Duration
		for {
			num_jellos := <-w.recv
			start = time.Now()
			for i := 0; i < num_jellos; i++ {
				jresult := w.gel.Jiggle(w.jello_recipe, w.jello_recipe)
				if jresult != w.jello_recipe+w.jello_recipe {
					panic("Customer found a bug in their Jello.")
				}
			}
			dur = time.Now().Sub(start)
			w.send <- dur
		}
	}()
}
