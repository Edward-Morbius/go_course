package main

import (
	"fmt"
)

type nType complex128

func (n nType) foo() int {
	return 11
}

func (n nType) bar() int {
	return 22
}

func (n nType) snafu() int {
	return 33
}

func main() {
	var fn func(nType) int
	fn = nType.snafu
	var y nType
	fmt.Println(fn(y))
}
