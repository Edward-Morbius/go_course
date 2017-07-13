package main

import (
	"fmt"
)

func main() {
	x := []byte{1, 23}
	y := []byte{1, 23}
	h := make(map[string]int)
	h[string(x)]++
	h[string(y)]++
	fmt.Println(h)
}
