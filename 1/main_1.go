package main

import "fmt"

type Human struct {
	name    string
	surname string
	age     int
}

func (h *Human) SayFullName() {
	fmt.Println(h.name + " " + h.surname)
}

func (h *Human) SayAge() {
	fmt.Println(h.age)
}

type Action struct {
	Human
}

func main() {
	act := Action{Human{"Ivan", "Ivanov", 22}}
	act.SayFullName()
	act.SayAge()
}
