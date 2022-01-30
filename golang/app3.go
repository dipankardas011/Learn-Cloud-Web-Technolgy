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

func switchStatements(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func switchStatementsv2() {
outerloop:
	for i := 0; i < 10; i++ {
		switch {
		case i&1 == 1:
			continue outerloop
		}
		fmt.Printf("[%v]\n", i)
	}
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
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
	switchStatements('x')
	switchStatementsv2()

}
