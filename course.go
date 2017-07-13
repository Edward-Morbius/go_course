package main

import (
	"fmt"
)

//myType is based on int
type myType int

//myType2 is based on myType is based on int!!!
type myType2 myType

type Never struct{}

func (mt myType) print() {
	fmt.Println(mt)
}

func main() {
	var x myType
	fmt.Println(myType(32) < 53)
	x = 99
	x.print()
}
