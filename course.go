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

func (c *counter) Read(b []byte) (int, error) {
	n, err := c.Reader.Read(b)
	c.n += n
	//Read reads n bytes into b and return the number of read bytes, therefore n not c.n.
	return n, err
}

func main() {
	inp := &counter{Reader: os.Stdin, n: 0}
	fmt.Println("␌")
	io.Copy(os.Stdout, inp)
	fmt.Println("␄", inp.n)
}
