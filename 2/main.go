package main

import (
	"fmt"
	"sync"
)

func square(val int) {
	fmt.Printf("%d^2 = %d\n", val, val*val)
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			square(val)
			wg.Done()
		}(val)
	}

	wg.Wait()
}
