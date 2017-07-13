package main

import (
	"fmt"
)

func main() {
	s := "H⇝ello"
	fmt.Println(s)
	fmt.Printf("%c\n", s[1])
	fmt.Println(len(s))
	//range assumes it is UTF-8!
	for i, c := range s {
		fmt.Printf("%2d: '%c'\n", i, c)
	}
	rs := []rune(s)
	fmt.Printf("%q\n", rs)
}
