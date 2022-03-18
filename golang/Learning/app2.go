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

func loops(N int) {
	var sum int = 0
	for i := 0; i < N; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)

	sum = 0

	//	while loop
	i := 0
	for i < N {
		sum += i
		i++
	}
	fmt.Println("Sum: ", sum)

	//	do while not there
	//	infinite loop
	for {
		if i == 100 {
			fmt.Println("Break cond.. ", i)
			return
		}
		i++
	}
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

	loops(10)
}
