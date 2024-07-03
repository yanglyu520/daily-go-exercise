package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		for i := 0; i < n; i++ {
			var a, b int
			_, err := fmt.Scanf("%d %d", &a, &b)
			if err != nil {
				break
			}
			fmt.Println(a + b)
		}

	}
}
