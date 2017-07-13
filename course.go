package main

import (
	"fmt"
)

//myType is based on int
type myType int

//myType2 is based on myType is based on int!!!
type myType2 myType

func main() {
	fmt.Println(myType(32) < myType2(53))
}
