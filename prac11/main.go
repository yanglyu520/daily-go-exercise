package main

import (
	"fmt"
	"reflect"
)

type E struct {
	Name string
}

func (e E) A() {
	fmt.Println("Method a")
}
func (e E) B() {
	fmt.Println("Method b")
}

func main() {
	a := E{Name: "a"}
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(), t.NumMethod())

	b := "random string b"
	v := reflect.ValueOf(b)
	v.SetString("new")
	fmt.Println(b, v)
}
