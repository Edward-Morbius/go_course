package main

import (
	"fmt"
)

//myType is based on int
type myType int

func main() {
	fmt.Println(myType(32) < 53)
}
