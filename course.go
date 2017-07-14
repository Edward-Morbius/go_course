package main

import (
	"fmt"
)

func halfAnswer() int {
	return 21
}

func answer() int {
	return halfAnswer() + halfAnswer()
}

func whichAnser(half bool) int {
	if half {
		return halfAnswer()
	}
	return answer()
}

func main() {
	fmt.Printf("Answer: %d\n", answer())
}
