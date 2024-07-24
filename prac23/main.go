package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	b := make([]byte, r.Size())

	n, _ := r.Read(b)

	fmt.Println(n)

	file, err := os.Open("./x.txt")
	if err != nil {
		return
	}
	file.Read()

}
