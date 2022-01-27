package main

import (
	"fmt"
	"math/cmplx"
)

var (
	flag     bool       = false
	PID      uint64     = 1<<64 - 1
	complexP complex128 = cmplx.Sqrt(-5 + 12i)
)

func moreOnLoops() {
	//sum := 0

	//for _, value := range array {
	//	sum += value
	//}
	for pos, char := range "日本\x80語" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

func main() {
	fmt.Printf("Type: %T value: %v\n", flag, flag)
	fmt.Printf("Type: %T value: %v\n", PID, PID)
	fmt.Printf("Type: %T value: %v\n", complexP, complexP)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	moreOnLoops()
}
