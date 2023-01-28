package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в
//1 или 0.

func set_bit(number *int64, bit int, sign int) {
	if sign == 0 { //  проверяем если бит нужно установить в 0
		*number = *number &^ (1 << bit) //  (x &^ y = x & ~y)  y это 1 свинутая на нужный бит --> делаем логическое НЕ у (т.е. нужный бит стал 0) и побитовое И с х
	} else {
		*number = *number | (1 << bit) //  сдвигаем 1 на нужный бит и делаем логическое ИЛИ с числом
	}
}

func print_number_binary(number int64) {
	for k := 63; k >= 0; k-- { //  итерируемся по всем битам числа
		if (number & (1 << k)) != 0 { //  проверяем если на бите k стоит 1
			fmt.Print("1") //  печатаем 1
		} else {
			fmt.Print("0") //  печатаем 0
		}
	}
	fmt.Print("\n")
}

func main() {
	var number int64  //  создаем переменную для числа
	var bit, sign int //  создаем переменные для номера бита и знака, на который будем менять бит

	fmt.Print("Enter number: ")        //  просим ввести число
	fmt.Scan(&number)                  // считываем число
	fmt.Print("Enter number bit: ")    //  просим ввести номер бита для изменения
	fmt.Scan(&bit)                     //  считываем бит
	fmt.Print("Enter sign (0 or 1): ") // просим ввести знак, на который будем устанавливать бит
	fmt.Scan(&sign)                    // считываем знак

	print_number_binary(number) //  печатем число в двоичном представлении до замены бита
	set_bit(&number, bit, sign) //  задаем контретный бит в числе нужным значением
	print_number_binary(number) //  печатем число в двоичном представлении после замены бита
}
