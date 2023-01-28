package main

import "strings"

//К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}

var justString string

func createHugeString(count int) string { //  реализуем вызываемую функцию
	sb := strings.Builder{}
	for i := 0; i < count; i++ { //  составляем строку из count элементов
		sb.WriteString("A")
	}
	return sb.String() //  возвращаем получившуюся строку
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 { // проверяем длину
		// justString = v[:100] - сломает руну если там utf-8
		justString = string([]rune(v)[:100])
	}
}

func main() {
	someFunc()
}
