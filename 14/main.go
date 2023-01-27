package main

import (
	"fmt"
	"reflect"
)

func option1(x interface{}) string {
	return fmt.Sprintf("%T", x)
}

func option2(x interface{}) string {
	var res string
	switch x.(type) {
	case int:
		res = "int:"
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
	return res
}

func option3(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}

func main() {
	//  создаем срез из переменных разного типа
	arr := []any{1, 1.3, "hello", make(chan int), func() {}, false}

	//  Способ 1. С помощью спецификатора %Т
	fmt.Println("Version 1")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option1(val))
	}

	//  Способ 2. Использование переключателя
	fmt.Println("\nVersion 2")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option2(val))
	}

	//  Способ 3. Использование reflection
	fmt.Println("\nVersion 3")
	for _, val := range arr {
		fmt.Println("Type", val, "is", option3(val))
	}
}
