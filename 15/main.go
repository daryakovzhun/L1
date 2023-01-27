package main

import "strings"

var justString string

func createHugeString(count int) string {
	sb := strings.Builder{}
	for i := 0; i < count; i++ {
		sb.WriteString("A")
	}
	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 { // проверка длины
		// justString = v[:100] - сломает руну если там utf-8
		justString = string([]rune(v)[:100])
	}
}

func main() {
	someFunc()
}
