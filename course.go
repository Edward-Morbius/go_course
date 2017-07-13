package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("~~~")
	}()
	//No output!
}
