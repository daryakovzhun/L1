package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseWords(str string) string {
	words := strings.Fields(str)           //  разделяем строку на слова и записываем в массив
	var reverse string                     //  создаем результирующую строку
	for i := len(words) - 1; i >= 0; i-- { //  итерируемся в обратном порядке по массиву с словами строки
		reverse += words[i] //  доюавляем в строку слово
		if i != 0 {         //  если это не последнее слово,  то добавляем разделители в виде пробелов
			reverse += " "
		}
	}
	return reverse //  возвращаем результирующую строку
}

func main() {
	fmt.Println("Enter str: ")                           //  просим ввести строку
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') //  считываем строку до \n включительно
	str = strings.Trim(str, "\n")                        //  обрезаем \n
	fmt.Print(reverseWords(str))                         //  выводим на экран строку, где слова в обратном порядке
}
