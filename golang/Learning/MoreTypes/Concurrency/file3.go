package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	ch <- 10

	fmt.Println(<-ch)
}
