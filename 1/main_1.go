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
	fmt.Println(h.name + " " + h.surname) // печатаем полное имя
}

func (h *Human) SayAge() {
	fmt.Println(h.age) // печатаем возраст
}

type Action struct {
	Human // наследуемся от Human, также можно (*Human) и (h Human)
}

func main() {
	act := Action{Human{"Ivan", "Ivanov", 22}} // создаем структуру Action, поле которого является структура Human
	act.SayFullName()                          // вызов метода, который печатаем полное имя
	act.SayAge()                               // вызов метода, который печатаем возраст
}
