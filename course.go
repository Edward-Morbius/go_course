package main

import (
	"fmt"
	"io"
	"os"
)

type counter struct {
	io.Reader
	n int
}


func main() {
	fmt.Println("␌")
	io.Copy(os.Stdout, os.Stdin)
}
