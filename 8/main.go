package main

import "fmt"

func set_bit(number *int64, bit int, sign int) {
	if sign == 0 {
		*number = *number &^ (1 << bit % 64)
	} else {
		*number = *number | (1 << bit % 64)
	}
}

func print_number_binary(number int64) {
	for k := 63; k >= 0; k-- {
		if (number & (1 << k)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
	fmt.Print("\n")
}

func main() {
	var number int64
	var bit, sign int

	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	fmt.Print("Enter number bit: ")
	fmt.Scan(&bit)
	fmt.Print("Enter sign (0 or 1): ")
	fmt.Scan(&sign)

	print_number_binary(number)
	set_bit(&number, bit, sign)
	print_number_binary(number)
}
