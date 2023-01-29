package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество.

func main() {
	//  задаем массив из строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}) //  создаем словарь
	for _, val := range arr {        //  итерируемся по элементам массива
		if _, ex := set[val]; !ex { //  проверяем наличие элемента в словаре
			set[val] = struct{}{} //  если такого элемента нет, то добавляем его
		}
	}

	fmt.Println(set) //  выводим получившееся множество из данного массива
}
