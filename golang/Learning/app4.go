package main

import (
	"fmt"
	"math"
	"time"
)

const (
	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func timming() {
	fmt.Println("time f()")

	today := time.Now().Weekday()
	unixTime := time.Now().UnixNano()

	fmt.Println(today)
	fmt.Println(unixTime)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is saturday")
	case today + 1:
		fmt.Println("Tomorrow is saturday")
	case today + 2:
		fmt.Println("In 2 days")
	case today - 1:
		fmt.Println("Yesterday was saturday")
	default:
		fmt.Println("too far")
	}
}

func deferringAMethod() {
	var num int32
	num = 5
	fmt.Println(num)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("This Statement is going to be evaluated at the last")

	fmt.Println("Second last")
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
	timming()

	deferringAMethod()
}
