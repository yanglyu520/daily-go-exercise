package main

import "fmt"

func main() {
	for {
		var a, b int
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil || (a == 0 && b == 0) {
			break
		}
		fmt.Println(a + b)
	}
}
