package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func isUniqueLetter(s string) bool {
	set := make(map[rune]struct{})       //  создаем словарь куда будем записывать уникальные встречающиеся символы в строке
	rune_s := []rune(strings.ToLower(s)) //  переводим в нижний регистр и преобразуем в руну
	for _, v := range rune_s {           //  итерируемся по руне
		if _, ex := set[v]; !ex { //  проверяем нахождение символа в словаре
			set[v] = struct{}{} //  добавляем в словарь
		} else {
			return false // возвращаем false, так как символ не уникальный
		}
	}
	return true // возвращаем true
}

func main() {
	fmt.Println("Enter str: ")                           //  просим ввести строку
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') //  считываем строку до \n включительно
	str = strings.Trim(str, "\n")                        //  обрезаем \n

	fmt.Println(isUniqueLetter(str)) //  выводим поверку строки на уникальность символом
}
