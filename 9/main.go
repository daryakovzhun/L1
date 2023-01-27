package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch1 := make(chan float64)
	ch2 := make(chan float64)

	go func() {
		for {
			num := rand.Float64()
			fmt.Println("Random num = ", num)
			ch1 <- num
		}
	}()

	go func() {
		for val := range ch1 {
			fmt.Print("Get num from ch1 = ", val, ", ", val, "^2 = ", val*val, "\n")
			ch2 <- val * val
		}
	}()

	for val := range ch2 {
		fmt.Println("Get num from ch2 = ", val)
	}
}
