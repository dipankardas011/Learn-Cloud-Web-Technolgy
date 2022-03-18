package main

import (
	"fmt"
	"os"
)

func sum(s []int, c chan int) {
	fmt.Printf("CHild PID: %g PPID: %g\n", os.Getpid(), os.Getppid())
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	// same PPID and PID so it is thread no new process
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Printf("PArent PID: %g PPID: %g\n", os.Getpid(), os.Getppid())

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
