package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.

func main() {
	var a_str, b_str string           //  объявляем переменные для считывания числа с консоли
	fmt.Print("Enter big number1 : ") // просим ввести первое число
	fmt.Scan(&a_str)                  //  считываем первое число
	fmt.Print("Enter big number2 : ") // просим ввести второе число
	fmt.Scan(&b_str)                  // считываем второе число

	a := new(big.Int)      //  создаем переменную типа big.Int для первого числа
	a.SetString(a_str, 10) //  задаем значение переменной, считанное из консоли

	b := new(big.Int)      //  создаем переменную типа big.Int для второго числа
	b.SetString(b_str, 10) //  задаем значение переменной, считанное из консоли

	sum, sub, mul, div := new(big.Int), new(big.Int), new(big.Int), new(big.Float)                   //  создаем переменные для операций сложения, вычитания, умножения и деления
	fmt.Println("Sum:", a, "+", b, "=", sum.Add(a, b))                                               //  выводим результат сложения (a+b)
	fmt.Println("Sub:", a, "-", b, "=", sub.Sub(a, b))                                               //  выводим результат вычитания (a-b)
	fmt.Println("Mul:", a, "*", b, "=", mul.Mul(a, b))                                               //  выводим результат умножения (a*b)
	fmt.Println("Div:", a, "/", b, "=", div.Quo(new(big.Float).SetInt(a), new(big.Float).SetInt(b))) //  выводим результат деления (a/b)
}
