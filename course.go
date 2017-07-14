package main

import (
	"fmt"
)

func halfAnswer() int {
	return 20
}

func answer() int {
	return halfAnswer() + halfAnswer() + halfAnswer()
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
