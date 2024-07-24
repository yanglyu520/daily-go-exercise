package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// type must be capital to be exported
type Class struct {
	Title    string    `json:"title" info:"ccc"`
	Students []Student `json:"students" info:"yyy"`
}

type Student struct {
	ID   int
	Name string
}

func findTag(st interface{}) {
	t := reflect.TypeOf(st).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagStr := t.Field(i).Tag.Get("info")
		fmt.Println(tagStr)
	}
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]Student, 0),
	}
	for i := 0; i < reflect.TypeOf(*c).NumField(); i++ {
		fmt.Println(reflect.TypeOf(*c).Field(i))
	}

	for i := 0; i < reflect.TypeOf(*c).NumMethod(); i++ {
		fmt.Println(reflect.TypeOf(*c).Method(i))
	}

	// marshal struct-->JSON
	data, _ := json.Marshal(c)
	fmt.Println(string(data))

	// unmarshal
	// decode
	str := `{"Title":"101","Students":null}`
	c1 := &Class{}
	json.Unmarshal([]byte(str), c1)
	fmt.Println(c1)

	findTag(c)
}
