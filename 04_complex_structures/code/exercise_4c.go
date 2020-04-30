package main

import (
	"fmt"
)

func average(args []float64) float64 {
	result := 0.0
	scores := append([]float64{2}, args...)
	for _, num := range scores {
		result += num
	}
	fmt.Println("result", args)
	return result / float64(len(args))
}

func main() {
	var i = []float64{1, 2, 3}
	fmt.Println(average(i))
}
