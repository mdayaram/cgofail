package worker

import (
	"github.com/mdayaram/cgofail/jello"
	"runtime"
	"time"
)

type Worker struct {
	recv chan int
	send chan time.Duration
	gel  jello.Jello
}

func New(recv chan int, send chan time.Duration, jello jello.Jello) *Worker {
	return &Worker{recv: recv, send: send, gel: jello}
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
			w.gel.Jiggle()
			dur = time.Now().Sub(start)
			w.send <- dur
		}
	}()
}
