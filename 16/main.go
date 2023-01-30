package main

import (
	"fmt"
	"sort"
)

//  Реализовать быструю сортировку массива (quicksort) встроенными методами
//  языка.

type Human struct {
	name    string
	surname string
	age     int
}

func main() {
	//  задаем неотсортированный массив
	arr := []int{376, 609, 22, 54, 1, 0, 321, 999, 2321, 545}
	fmt.Println(arr) //  напечатаем до сортировки
	sort.Ints(arr)
	fmt.Println(arr) //  напетаем после сортировки

	humanArr := []Human{{"Petr", "Petrov", 50}, {"Ivan", "Ivanov", 23}, {"Kate", "Orlova", 45}}
	fmt.Println(humanArr)
	sort.Slice(humanArr, func(i, j int) bool {
		return humanArr[i].surname < humanArr[j].surname
	})
	fmt.Println(humanArr)
}
