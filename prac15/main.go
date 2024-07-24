package main

import "fmt"

func main() {
	adder := Closure()
	fmt.Println(adder(1)) // Output: 1
	fmt.Println(adder(2)) // Output: 3
	fmt.Println(adder(3)) // Output: 6
}

func Closure() func(int) int {
	sum := 0
	return func(y int) int {
		sum += y
		return sum
	}
}
