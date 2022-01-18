package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {

	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 5)

	fmt.Println(arr)

	fmt.Println(add(3, 5))
	first := "First"
	second := "Second"

	a, b := swap(first, second)
	fmt.Println(a, b)
}
