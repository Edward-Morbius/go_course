package main

import (
	"fmt"
)

func print(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println("int", x)
	case float32:
		fmt.Println("int", x)
	case string:
		fmt.Println("string", x)
	default:
		fmt.Printf("unknown type `%s`: '%s'\n", t, x)
	}
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
