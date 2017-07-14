package main

import (
	"fmt"
	"time"
)

func main() {
	var problem int
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				problem++
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(problem)
}
