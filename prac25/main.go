package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("human.eat")
}

type Teacher struct {
	Human   // inherit all functions human class has
	subject string
}

// =========================
func (t *Teacher) Eat() {
	fmt.Printf("teacher can eat after teaching %s\n", t.subject)
}

func (t *Teacher) Teach() {
	fmt.Printf("teacher can teach %s\n", t.subject)
}
func main() {
	h := Human{
		name: "yang",
		sex:  "female",
	}

	t := Teacher{
		Human:   h,
		subject: "math",
	}

	t.Eat()
}
