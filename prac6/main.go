package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal(err)
	}

	var x int
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = x
	}

	for j := n - 1; j >= 0; j-- {
		fmt.Printf("%d ", nums[j])
	}

	fmt.Println()
	for k := 1; k < n; k = k + 2 {
		fmt.Printf("%d ", nums[k])
	}

}
