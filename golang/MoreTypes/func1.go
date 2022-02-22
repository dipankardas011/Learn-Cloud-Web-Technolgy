package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, variable int32) float64 {
	fmt.Println("val: ", variable)
	return fn(float64(variable), 0)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	summer := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(summer(5, 12))
	fmt.Println(compute(summer, 5))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
