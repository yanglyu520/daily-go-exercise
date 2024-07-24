package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	name string
	age  int
}
type teacher struct {
	classroom string
	person
}

func main() {
	p := newPerson("yang", 1)
	fmt.Println(*p)
	fmt.Println(p.GetName())

	tea := teacher{
		classroom: "302",
	}
	fmt.Println(tea.classroom, tea.name, tea.age)
}

// 结构体的实例化
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 定义方法，p是
func (p person) GetName() string {
	return p.name
}
