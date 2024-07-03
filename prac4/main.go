package main

import "fmt"

func main() {
	var n int // number of lines
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		for i := 0; i < n; i++ {
			var m int
			_, err := fmt.Scanf("%d", &m)
			if err != nil {
				break
			}

			sum := 0
			for j := 0; j < m; j++ {
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
}
