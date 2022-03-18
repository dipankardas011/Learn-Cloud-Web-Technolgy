package main

import (
	"fmt"
	"strings"
)

func slices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func ranger() {
	var pow = []int{1, 2, 4, 8}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// by skipping the index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// this will print the index
	for i := range pow {
		fmt.Println(i)
	}

	//	more range use
	powee := make([]int, 10)
	for i := range powee {
		powee[i] = 1 << i
	}

	for _, value := range powee {
		fmt.Println(value)
	}
}

func main() {
	slices()
	ranger()
}
