package main

import "fmt"

type Human3 struct {
	name    string
	surname string
	age     int
}

func (h *Human3) SayFullName() {
	fmt.Println(h.name + " " + h.surname)
}

func (h *Human3) SayAge() {
	fmt.Println(h.age)
}

type Action3 struct {
	h Human3
}

func main() {
	act := Action3{Human3{"Ivan", "Ivanov", 22}}
	act.h.SayFullName()
	act.h.SayAge()
}
