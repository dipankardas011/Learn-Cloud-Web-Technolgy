package main

import "fmt"

func fff() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("NILLL")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func gdgd() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
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
	gdgd()
}
