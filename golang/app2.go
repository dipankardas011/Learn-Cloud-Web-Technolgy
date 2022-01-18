package main

import (
	"fmt"
)

func namedRet(sum int) (x, y int) {
	x = sum * 4
	y = sum + 1

	return
}

func print(x bool) {
	first, last := "Dipankar", "Das"
	fmt.Println("Name:", first, last, x)
}

func main() {
	var aa int
	fmt.Println(aa)

	var i, j int
	i = 0
	j = 0
	fmt.Println(i, j)

	var name string = "Dipankar Das"
	fmt.Println(name)

	var res bool
	fmt.Println(res)

	print(true)

	fmt.Println(namedRet(15))
}
