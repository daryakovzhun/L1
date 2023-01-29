package main

import (
	"fmt"
	"log"
	"time"
)

//Реализовать пересечение двух неупорядоченных множеств.

func intersection(arr1, arr2 []int) []int {
	defer duration(track("intersection1")) //  вызываем функцию определения времени выполнения
	var res []int                          //  создаем результирующее множество
	for _, val1 := range arr1 {            //  итерируемся по 1 множеству
		for _, val2 := range arr2 { //  итерируемся по 2 множеству
			if val1 == val2 { //  проверяем совпадения значений из множеств
				res = append(res, val1) //  если значения сомпадают добавляем в результирующее множество и выходим из внутреннего цикла
				break
			}
		}
	}
	return res //  возвращаем результирующее множество
}

func intersection2(arr1, arr2 []int) []int {
	defer duration(track("intersection2")) //  вызываем функцию определения времени выполнения
	set := make(map[int]struct{})          //  создаем словарь, который будет выступать в роли множества

	small, big := func() ([]int, []int) { //  определяем большое и маленькое множество
		if len(arr1) <= len(arr2) {
			return arr1, arr2
		} else {
			return arr2, arr1
		}
	}()

	for _, val1 := range small { //  итерируемся по меньшему множеству и заполняем словарь
		set[val1] = struct{}{}
	}

	var res []int              //  создаем результирующее множество
	for _, val2 := range big { //  итерируемся по множеству большей длины
		if _, ex := set[val2]; ex { //  проверяем наличие элемента из множества в словаре
			res = append(res, val2) //  если такое значение есть, то добавляем его в результирующее множество
		}
	}

	return res //  возвращаем результирующее множество
}

func main() {
	//  задаем 2 массива, выполняющие роль множеств (уникальные эементы)
	arr1 := []int{66, 1234, 5432, 4, 3, 1, 6, 8, 34, 99, 5, 876, 34, 88, 22, 95, 67, 44, 11, 875, 243, 456, 90, 8676, 345, 43434, 6343}
	arr2 := []int{23, 65, 5, 4, 78, 1}
	result1 := intersection(arr1, arr2)  //  ищем пересечение 1 способом
	result2 := intersection2(arr1, arr2) //  ищем пересечение 2 способом

	fmt.Println(result1) //  выводим результат
	fmt.Println(result2) //  выводим результат
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// функция для определения времени выполнения, вызываемой функции
func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
