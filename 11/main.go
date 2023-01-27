package main

import (
	"fmt"
	"log"
	"time"
)

func intersection(arr1, arr2 []int) []int {
	defer duration(track("intersection1"))
	var res []int
	for _, val1 := range arr1 {
		for _, val2 := range arr2 {
			if val1 == val2 {
				res = append(res, val1)
				break
			}
		}
	}
	return res
}

func intersection2(arr1, arr2 []int) []int {
	defer duration(track("intersection2"))
	set := make(map[int]bool)

	small, big := func() ([]int, []int) {
		if len(arr1) <= len(arr2) {
			return arr1, arr2
		} else {
			return arr2, arr1
		}
	}()

	for _, val1 := range small {
		set[val1] = false
	}

	for _, val2 := range big {
		if _, ex := set[val2]; ex {
			set[val2] = true
		}
	}

	var res []int
	for k, v := range set {
		if v {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	arr1 := []int{66, 1234, 5432, 4, 3, 1, 6, 8, 34, 99, 5, 876, 34, 88, 22, 95, 67, 44, 11, 875, 243, 456, 90, 8676, 345, 43434, 6343}
	arr2 := []int{23, 65, 5, 4, 78, 1}
	result1 := intersection(arr1, arr2)
	result2 := intersection2(arr1, arr2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
