package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, val := range arr {
		if _, ex := set[val]; !ex {
			set[val] = true
		}
	}

	fmt.Println(set)
}
