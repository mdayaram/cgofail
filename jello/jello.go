package jello

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* jiggle(char* one, char* two) {
	int i, lenOne, lenTwo, sizeOne, sizeTwo;
	char* newbuff;

	if ( one == NULL && two == NULL ) {
		return "";
	} else if ( one == NULL ) {
		return two;
	} else if ( two == NULL ) {
		return one;
	}

	lenOne = strlen(one);
	lenTwo = strlen(two);
	sizeOne = sizeof(one) * lenOne;
	sizeTwo = sizeof(two) * lenTwo;
	newbuff = (char*) malloc(sizeOne + sizeTwo + 1);
	for(i = 0; i < lenOne; i++) {
		newbuff[i] = one[i];
	}
	for(i = 0; i < lenTwo; i++) {
		newbuff[lenOne + i] = two[i];
	}
	newbuff[lenOne + lenTwo] = '\0';
	return newbuff;
}
*/
import "C"
import "unsafe"

type Jello interface {
	Jiggle(string, string) string
}

type Cgo struct{}
type Gor struct{}

func NewCgo() *Cgo {
	return &Cgo{}
}

func (c *Cgo) Jiggle(one, two string) string {
	cone := C.CString(one)
	defer C.free(unsafe.Pointer(cone))
	ctwo := C.CString(two)
	defer C.free(unsafe.Pointer(ctwo))

	cresult := C.jiggle(cone, ctwo)
	result := C.GoString(cresult)
	C.free(unsafe.Pointer(cresult))
	return result
}

func NewGor() *Gor {
	return &Gor{}
}

func (g *Gor) Jiggle(one, two string) string {
	newbuf := one + two
	return newbuf
}
