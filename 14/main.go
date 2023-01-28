package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

func option1(x interface{}) string {
	return fmt.Sprintf("%T", x) //  возврщаем тип переменной с помощью спеифкатора %Т
}

func option2(x interface{}) string {
	var res string
	switch x.(type) { //  определяем тип интерфейса с помощью .(type) и проверяем на соответсвие
	case int:
		res = "int:" // записываем в строку тип перемнной
	case float64:
		res = "float64:"
	case bool:
		res = "bool"
	case string:
		res = "string"
	case chan int:
		res = "chan int"
	case func():
		res = "func()"
	default:
		fmt.Println("unknown")
	}
	return res //  возвращаем строку
}

func option3(x interface{}) reflect.Type {
	return reflect.TypeOf(x) //  возвращаем тип с помощью reflect.TypeOf()
}

func main() {
	//  создаем срез из переменных разного типа
	arr := []any{1, 1.3, "hello", make(chan int), func() {}, false}

	//  Способ 1. С помощью спецификатора %Т
	fmt.Println("Version 1")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option1(val)) //  выводим на экран значение и тип переменной, полученной функцией option1
	}

	//  Способ 2. Использование переключателя
	fmt.Println("\nVersion 2")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option2(val)) //  выводим на экран значение и тип переменной, полученной функцией option2
	}

	//  Способ 3. Использование reflection
	fmt.Println("\nVersion 3")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option3(val)) //  выводим на экран значение и тип переменной, полученной функцией option3
	}
}
