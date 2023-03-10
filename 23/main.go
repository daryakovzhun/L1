package main

import "fmt"

//Удалить i-ый элемент из слайса.

func removeQuik(arr []int, idx int) []int {
	arr[idx] = arr[len(arr)-1] //  копируем последний элемент в индекс idx
	arr[len(arr)-1] = 0        //  удаляем последний элемент (записываем нулевое значение)
	arr = arr[:len(arr)-1]     // усекаем срез
	return arr
}

func remove(arr []int, idx int) []int {
	copy(arr[idx:], arr[idx+1:]) //  выполняем сдвиг arr[i+1:] влево на один индекс
	arr[len(arr)-1] = 0          //  удаляем последний элемент (записываем нулевое значение)
	arr = arr[:len(arr)-1]       //  усекаем срез
	return arr
}

func main() {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var index int              //  переменная для индекса
	fmt.Print("Enter index: ") //  просим ввести индекс для удаления элемента по этому индексу
	fmt.Scan(&index)           //  считываем индекс

	arr1 = removeQuik(arr1, index)                 // работает быстрее, но меняет порядок элементов
	arr2 = remove(arr2, index)                     // работает медленее, но сохраняет порядок
	arr3 = append(arr3[:index], arr3[index+1:]...) // встроенный метод

	fmt.Println(arr1) //  выводим получившийся слайс
	fmt.Println(arr2) //  выводим получившийся слайс
	fmt.Println(arr3) //  выводим получившийся слайс
}
