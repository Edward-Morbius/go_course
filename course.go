package main

import (
	"fmt"
)

type base int

type sub struct {
	base
}

func (b base) print() {
	fmt.Println("base", b)
}

//Overload for derived type
func (s sub) print() {
	s.base.print()
	fmt.Println("sub", s)
}

/*
 Only available for methods, no polymorphism for functions.

 This is an error:
*/
// func fun(x int) {
// 	fmt.Println("fun, int")
// }

// func fun(float32 f) {
// 	fmt.Println("fun, float")
// }

func main() {
	var b base = 9
	b.print()
	var s sub
	s.print()
}
