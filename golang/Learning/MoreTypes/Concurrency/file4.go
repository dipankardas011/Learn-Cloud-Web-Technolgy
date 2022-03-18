package main

import "fmt"

func fabonacii(n int, pipe chan int) {
	a, b := -1, 1
	for i := 0; i < n; i++ {
		c := a + b
		pipe <- c
		a = b
		b = c
	}
	close(pipe)
}

func main() {
	channel := make(chan int, 10)
	fmt.Println("Enter the value for N")
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}

	go fabonacii(N, channel)
	for i := range channel {
		fmt.Print(" ", i)
	}
	fmt.Println()
}
