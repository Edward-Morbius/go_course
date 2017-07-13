package main

import (
	"fmt"
)

//Interface definition.
type vehicle interface {
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
}
