package main

import "fmt"

func fff() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("NILLL")
	}
}

func main() {
	users := []struct {
		name   string
		age    int8
		gender string
	}{
		{"Dipankar", 23, "M"},
		{"HEli", 21, "F"},
	}

	fmt.Println(users)
	fff()
}
