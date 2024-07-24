package main

import "fmt"

func main() {
	func() {
		fmt.Println("some func")
	}()

	func(name string) {
		fmt.Println("name is", name)
	}("uamg")

	name := func() string {
		return "jiajiajia"
	}() // name is not a func val but a val for the returned val
	fmt.Println(name)
}
