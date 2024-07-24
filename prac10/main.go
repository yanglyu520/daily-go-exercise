package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var rw io.ReadWriter
	f, _ := os.Open("test.txt")
	rw = f

	x, ok := rw.(*os.File)
	fmt.Println(rw, x, ok)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(rw))
}
