package main

import "fmt"

type Human2 struct {
	name    string
	surname string
	age     int
}

func (h *Human2) SayFullName() {
	fmt.Println(h.name + " " + h.surname)
}

func (h *Human2) SayAge() {
	fmt.Println(h.age)
}

type Action2 struct {
	*Human2
}

func main() {
	human := Human2{"Ivan", "Ivanov", 22}
	act := Action2{&human}
	act.SayFullName()
	act.SayAge()
}
