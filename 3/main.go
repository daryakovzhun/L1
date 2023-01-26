package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func option1() {
	arr := []int{2, 4, 6, 8, 10}
	chan_out := make(chan int, len(arr))
	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			chan_out <- val * val
			wg.Done()
		}(val)
	}

	wg.Wait()
	close(chan_out)

	var sum int
	for v := range chan_out {
		sum += v
	}

	fmt.Println("Sum1 = ", sum)
}

func option2() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	var sum int32
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			atomic.AddInt32(&sum, int32(val*val))
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("Sum2 = ", sum)
}

func option3() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	var sum int
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			mu.Lock()
			sum += val * val
			mu.Unlock()
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("Sum3 = ", sum)
}

func main() {
	option1()
	option2()
	option3()
}
