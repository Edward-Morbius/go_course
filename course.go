package main

import (
	"bytes"
	"fmt"
	"strings"
)

func repeat1(s string, n int) string {
	var x string
	for i := 0; i < n; i++ {
		x += s
	}
	return x
}

func repeat2(s string, n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}


func repeat3(s string, n int) string {
	return strings.Repeat(s, n)
}

func main() {
	s := repeat1("Hello, World!\n", 10)
	fmt.Println(s)
}
