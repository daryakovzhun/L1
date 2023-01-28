package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

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
