package main

import "fmt"

func main() {
	var n int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		sum := 0
		for i := 0; i < n; i++ {
			var x int
			_, err := fmt.Scanf("%d", &x)
			if err != nil {
				break
			}
			sum += x
		}
		fmt.Printf("%d\n", sum)

	}
}
