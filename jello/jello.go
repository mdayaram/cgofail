package jello

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* jiggle(char* one, char* two) {
	int i, lenOne, sizeOne, sizeTwo;
	char* newbuff;

	if ( one == NULL && two == NULL ) {
		return "";
	} else if ( one == NULL ) {
		return two;
	} else if ( two == NULL ) {
		return one;
	}

	lenOne = strlen(one);
	sizeOne = sizeof(one) * lenOne;
	sizeTwo = sizeof(two) * strlen(two);
	newbuff = (char*) malloc(sizeOne + sizeTwo + 1);
	for(i = 0; i < strlen(one); i++) {
		newbuff[i] = one[i];
	}
	for(i = 0; i < strlen(two); i++) {
		newbuff[lenOne + i] = two[i];
	}
	newbuff[lenOne + strlen(two)] = '\0';
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
	newbuf := ""
	for _, r := range one {
		newbuf += string(r)
	}
	for _, r := range two {
		newbuf += string(r)
	}
	return newbuf
}
