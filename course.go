package main

import (
	"fmt"
)

func main() {
	x := []byte{1, 23}
	y := []byte{1, 23}
	h := make(map[string]int)
	/*
		h[x]++

		# github.com/Edward-Morbius/go_course
		./course.go:11: cannot use x (type []byte) as type string in map index

	*/

	h[string(y)]++
	fmt.Println(h)
}
