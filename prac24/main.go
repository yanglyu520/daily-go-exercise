package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx2 := context.TODO()
	ctx3, _ := context.WithCancel(ctx)
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx2)
	fmt.Println(ctx3)
}
