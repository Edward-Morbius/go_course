package main

import (
	"fmt"
	"time"
)

func main() {
	//Capacity of channel is two: ch := make(chan int, 2)
	ch := make(chan int)
	go func() {
		x := <- ch
		fmt.Println("~~~", x)
	}()
	ch <- 100
	time.Sleep(time.Second)
}
