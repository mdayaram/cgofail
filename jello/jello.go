package jello

/*
#include <stdio.h>
#include <stdlib.h>

int jiggle() {
	int i, sum = 0;
	for(i = 0; i < 10000; i++) {
		sum += i;
	}
	return sum;
}
*/
import "C"

type Jello interface {
	Jiggle() int
}

type Cgo struct{}
type Gor struct{}

func NewCgo() *Cgo {
	return &Cgo{}
}

func (c *Cgo) Jiggle() int {
	return int(C.jiggle())
}

func NewGor() *Gor {
	return &Gor{}
}

func (g *Gor) Jiggle() int {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	return sum
}
