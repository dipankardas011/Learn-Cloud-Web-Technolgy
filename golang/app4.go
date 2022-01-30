package main

import (
	"fmt"
	"math"
)

const (
	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var i int = 43
	var f float64 = float64(i)
	var u uint = uint(f)
	x := 4
	y := 3

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
	res := math.Sqrt(float64(x*x + y*y))
	fmt.Printf("%0.2f type %T\n", res, res)

	const PIE = math.Pi
	fmt.Printf("%2.5f %T\n", PIE, PIE)
	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(SMALL))
}
