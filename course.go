package main

import (
	"fmt"
)

//myType is based on int
type myType int

//myType2 is based on myType is based on int!!!
type myType2 myType

type Never struct{}

func main() {
	fmt.Println(myType(32) < 53)
}
