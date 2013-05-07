package worker

/*
#include <stdio.h>
#include <stdlib.h>

int jiggle() {
	int i, sum = 0;
	for(i = 0; i < 10000; i++) {
		sum += i;
	}
	return i;
}
*/
import "C"

import (
	"time"
)

type Worker struct {
	recv chan int
	send chan time.Duration
}

func New(recv chan int, send chan time.Duration) *Worker {
	return &Worker{recv: recv, send: send}
}

func (w *Worker) WorkIt() {
	go func() {
		var start time.Time
		var dur time.Duration
		for {
			<-w.recv
			start = time.Now()
			C.jiggle()
			dur = time.Now().Sub(start)
			w.send <- dur
		}
	}()
}
