package main

import "fmt"

func pattern1() {
	n := 5
	for i := 0; i < n; i++ {
		for sp := n; sp >= i; sp-- {
			fmt.Print(" ")
		}
		for j := i; j < 2*i; j++ {
			fmt.Print(j)
		}
		for k := 2*i - 2; k >= i; k-- {
			fmt.Print(k)
		}
		fmt.Println()
	}
}

func main() {
	pattern1()
}
