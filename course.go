package main

import (
	"fmt"
)

func slowRepeat(s string, n int) string {
	var x string
	for i := 0; i < n; i++ {
		x += s
	}
	return x
}

func main() {
	s := slowRepeat("Hello, World!\n", 10)
	fmt.Println(s)
}
