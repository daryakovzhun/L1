package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverse(s *string) {
	rune_s := []rune(*s)                                   //  преобразуем строку в руну
	for i, j := 0, len(rune_s)-1; i < j; i, j = i+1, j-1 { //  заводим 2 счетчика с начала и конца строки
		rune_s[i], rune_s[j] = rune_s[j], rune_s[i] //  меняем местами первый и последний элемент и т.д.
	}
	*s = string(rune_s) //  преобразыем обратно в строку
}

func main() {
	fmt.Print("Enter str: ")                             //  просим ввести строку
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') //  считываем строку до \n включительно
	str = strings.Trim(str, "\n")                        //  обрезаем \n

	reverse(&str)    //  вызываем функцию переворота строки
	fmt.Println(str) // выводим получившуюся строку
}
