package main

import "strings"

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
