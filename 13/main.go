package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 12
	b := 5
	fmt.Println("BEFORE: a =", a, "b =", b)

	//  параллельное присваивание
	a, b = b, a
	fmt.Println("AFTER (||): a =", a, "b =", b)

	//  сложение/вычитание
	a = a + b //  a = 12 + 5 = 17
	b = a - b //  b = 17 - 5 = 12
	a = a - b //  a = 17 - 12 = 5
	fmt.Println("AFTER (+/-): a =", a, "b =", b)

	//  умножение/деление
	a = a * b //  a = 12 * 5 = 60
	b = a / b //  b = 60 / 5 = 12
	a = a / b //  a = 60 / 12 = 5
	fmt.Println("AFTER (*/): a =", a, "b =", b)
	//Главным недостатком является большее количество операций,
	//в чём можно убедиться посчитав операции сложения, вычитания и присваивания.
	//Второй важный нeдостаток это область применения — числа. Арифметика для
	//вeщeствeнных чисeл можeт выполняться некорректно, что приведёт к неожиданному результату.

	//  битовые операции XOR
	//  a = 1100 b = 0101
	a = a ^ b //  1100 XOR 0101 = 1001 (9)
	b = b ^ a //  0101 XOR 1001 = 1100 (12)
	a = a ^ b //  1001 XOR 1100 = 0101 (5)
	fmt.Println("AFTER (^=): a =", a, "b =", b)
}
