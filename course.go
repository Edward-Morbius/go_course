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

func main() {
	var b base = 9
	b.print()
	var s sub
	s.print()
}
