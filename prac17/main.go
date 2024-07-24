package main

import "fmt"

func create() func() int {
	c := 2 // captured
	return func() int {
		c += 1
		return c
	}
}

func main() {
	f1 := create()
	f2 := create()

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f2())
}
