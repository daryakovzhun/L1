package main

import (
	"fmt"
	"math"
)

func main() {
	//  задаем последовательность температурных колебаний
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := make(map[int][]float64) //  создаем результирующий словарь

	for _, val := range arr { //  итерируемся по элементам массива
		if val > 0 { //  если значение элемента массива больше 0, то округляем в меньшую сторону (функция Floor)
			result[int(math.Floor(val/10)*10)] = append(result[int(math.Floor(val/10)*10)], val) //  добавляем в список словаря по заданному ключу
		} else { //  в противном случае округляем большую сторону
			result[int(math.Ceil(val/10)*10)] = append(result[int(math.Ceil(val/10)*10)], val)
		}
	}

	fmt.Println(result) //  выводим получившиеся группы
}
