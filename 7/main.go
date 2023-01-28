package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

type Counters struct {
	mx sync.Mutex
	m  map[int]int
}

func NewCounters() *Counters {
	return &Counters{
		m: make(map[int]int),
	}
}

func (c *Counters) Store(key int, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func main() {
	arr := []int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}

	fmt.Println("sync.Mutex")
	// использование sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(1)
	counters := NewCounters()
	go func() {
		for k, v := range arr {
			counters.Store(k, v)
		}
		wg.Done()
	}()
	wg.Wait()

	for k, v := range counters.m {
		fmt.Println("Key = ", k, " , value = ", v)
	}

	fmt.Println("\nsync.Map")
	// использование sync.Map
	var counters1 sync.Map
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		for k, v := range arr {
			counters1.Store(k, v)
		}
		wg1.Done()
	}()
	wg1.Wait()

	counters1.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true // if false, Range stops
	})
}
