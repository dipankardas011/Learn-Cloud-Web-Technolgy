package main

import "fmt"

func hhh() {
	hello := []string{
		"aaaa",
		"aaaasss",
		"dcdsc",
		"aaaasvfv fdfsss",
		"aaaadfvfdvsss",
		"aaaav   sss",
	}

	fmt.Println(hello)
	slice := hello[1:4]
	fmt.Println(slice)
	slice[0] = "DDDDIDIDIDI"
	fmt.Println(hello)
}

func main() {
	var arr1 [10]string

	arr1[0] = "aaaa"
	arr1[1] = "dfscsd"
	fmt.Println(arr1)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	var slicting []int = primes[1:4]
	fmt.Println(slicting)

	hhh()
}
