package main

/*
#include <stdlib.h>
*/
import "C"

//export GoAdd
func GoAdd(lhs, rhs int) int {
	return lhs + rhs
}

func main() {}
