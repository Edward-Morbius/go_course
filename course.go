package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type node struct {
	next *node
	v    int
}

func create(n int) *node {

	if n < 1 {
		return nil
	}

	head := &node{v: 0}

	current := head

	for i := 1; i < n; i++ {
		n := &node{v: i}
		current.next = n
		current = n
	}

	return head
}

func (n *node) len() int {
	var l int
	for ; n != nil; n, l = n.next, l+1 {
	}
	return l
}

func main() {

	n := flag.Int("n", 1e6, "number of nodes in list")
	c := flag.String("c", "", "write cpu profile to file")

	flag.Parse()

	if *c != "" {
		f, err := os.Create(*c)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	list := create(*n)
	fmt.Printf("%d\n", list.len())
}
