package main

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
}

func (a A) Name() string {
	return "hi, " + a.name
}
func Name2(a A) string {
	return "hi, " + a.name
}
func main() {
	a := A{name: "x"}
	fmt.Println(a.Name())
	fmt.Println(A.Name(a))
	t1 := reflect.TypeOf(A.Name)
	t2 := reflect.TypeOf(Name2(a))
	fmt.Println(t1 == t2)
}
