package main

import (
	"fmt"
)

func goRoutine(ch <-chan int, done chan struct{}) {
	defer close(done)
	for x := range ch {
		fmt.Println("⇝⇝⇝", x)
	}
	fmt.Println("⇜⇜⇜")
}

func main() {
	//Capacity of channel is two: ch := make(chan int, 2)
	ch := make(chan int)
	done := make(chan struct{})
	go goRoutine(ch, done)
	ch <- 100
	ch <- 101
	ch <- 102
	close(ch)
	<-done
}
