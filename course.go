package main

import (
	"fmt"
)

func fn(slice []byte) [4]byte {
	var arr [4]byte
	copy(arr[:], slice)
	return arr
}

func main() {
	x := []byte{1, 23, 7, 42}
	h := make(map[[4]byte]int)
	y := fn(x)
	h[y] = 111
	fmt.Println(h)
}
