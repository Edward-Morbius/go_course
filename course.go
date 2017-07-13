package main

import (
	"fmt"
)

func deepNested() int {
	panic("ERROR")
	fmt.Println("Nested")
	return 42
}

func intensiveAnalysis() (y int, err error) {
	defer func() {
		if x:= recover(); x != nil {
			err = fmt.Errorf("Failed %#v", x)
		}
	}()

	//
	x := deepNested()
	return x, nil
}


func main() {
	val, err := intensiveAnalysis()
	fmt.Println(val, err)
}
