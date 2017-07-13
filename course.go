package main

import (
	"fmt"
)

//myType is based on int
type myType int

//myType2 is based on myType is based on int!!!
type myType2 myType

type Never struct{}

//Remember that mt is BY VALUE!
//    ↓↓↓↓↓↓↓↓↓↓ Receiver
func (mt myType) print() {
	fmt.Println(mt)
}

//    ↓↓↓↓↓↓↓↓↓↓↓ Pointer receiver.
func (mt *myType) modify() {
	if mt != nil {
		*mt = 42
	}
}

func main() {
	var x myType
	fmt.Println(myType(32) < 53)
	x = 99
	x.print()
	x.modify()
	x.print()
	var y *myType
	//This works!
	y.modify()
	//This fails!
	y.print()
}
