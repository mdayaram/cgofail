package worker

import (
	"github.com/mdayaram/cgofail/jello"
	"runtime"
	"time"
	"io/ioutil"
)

type Worker struct {
	recv chan int
	send chan time.Duration
	gel  jello.Jello
	jiggles int
	text string
}

func New(recv chan int, send chan time.Duration, jello jello.Jello, jiggles int) *Worker {
	inbytes, err := ioutil.ReadFile("lol.txt")
	if err != nil {
		panic("FILE DAMMIT!")
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
				if jresult != w.text + w.text {
					panic("DAMMIT")
				}
			}
			dur = time.Now().Sub(start)
			w.send <- dur
		}
	}()
}
