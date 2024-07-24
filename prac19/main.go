package main

import (
	"fmt"
	"unsafe"
)

type Myint int
type MyAliasInt = int
type Book struct {
	name string
}

type Foo struct {
	d struct{}
	a int8  // 1 byte
	c bool  // 1 byte
	b *Book // 8 byte pointer type will be depend

}

func main() {
	//var a Myint
	//var b MyAliasInt
	//
	//fmt.Printf("a type: %T, b type: %T\n", a, b)
	//
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(b))

	//var c Book
	//d := new(Book)
	//fmt.Println(c)
	//fmt.Println(d)
	//c.name = "ssss"
	//fmt.Println(c)
	//d.name = "ooo"
	//fmt.Println(d)
	var g struct{}
	var f Foo
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Alignof(f.a))
	fmt.Println(unsafe.Alignof(f.c))
	fmt.Println(unsafe.Alignof(f.b))
	fmt.Println(unsafe.Alignof(f.d))
	fmt.Println(unsafe.Sizeof(g))
}
