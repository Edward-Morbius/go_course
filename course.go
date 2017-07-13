package main

import (
	"fmt"
)

func print(x interface{}) {
}


//Interface definition.
type vehicle interface {
	//Only methods!
	start()
	stop()
}

type car int

func (c car) start() {
}

func (c car) stop() {
}

//Now magically car implements the interface of vehicle as it has all
//the functions...

func do(v vehicle) {
}

//Does not wirk with nil
type something struct{}

func main() {
	var c car

	fmt.Println(c)
	do(c)

	//Multitype, variadic functions
	print(1)
	print(1.1111)
	print("Hello")
}
