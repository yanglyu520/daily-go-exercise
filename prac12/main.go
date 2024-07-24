package main

import "fmt"

func demo(name string, hover ...string) {
	fmt.Println(name)
	for i, n := range hover {
		fmt.Println(i, n)
	}
}

func main() {
	demo("yang love", "food", "yoga", "drawing")
	demo("yang")
	a := []int{1, 2, 3}
	b := append([]int{}, a...)
	fmt.Println(a, b)
}
