package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := make(map[int][]float64)

	for _, val := range arr {
		if val > 0 {
			result[int(math.Floor(val/10)*10)] = append(result[int(math.Floor(val/10)*10)], val)
		} else {
			result[int(math.Ceil(val/10)*10)] = append(result[int(math.Ceil(val/10)*10)], val)
		}
	}

	fmt.Println(result)
}
