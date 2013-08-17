package worker

import (
	"io/ioutil"
	"runtime"
	"time"

	"github.com/mdayaram/cgofail/jello"
)

type Worker struct {
	recv    chan int
	send    chan time.Duration
	gel     jello.Jello
	jiggles int
	text    string
}

func New(recv chan int, send chan time.Duration, jello jello.Jello, jiggles int, recipe string) *Worker {
	inbytes, err := ioutil.ReadFile(recipe)
	if err != nil {
		panic(err)
	}
	instr := string(inbytes)
	return &Worker{recv: recv, send: send, gel: jello, jiggles: jiggles, text: instr}
}

func (w *Worker) WorkIt(lockThread bool) {
	go func() {
		if lockThread {
			runtime.LockOSThread()
		}

		var start time.Time
		var dur time.Duration
		for {
			<-w.recv
			start = time.Now()
			for i := 0; i < w.jiggles; i++ {
				jresult := w.gel.Jiggle(w.text, w.text)
				if jresult != w.text+w.text {
					panic("Customer found a bug in their Jello.")
				}
			}
			dur = time.Now().Sub(start)
			w.send <- dur
		}
	}()
}
