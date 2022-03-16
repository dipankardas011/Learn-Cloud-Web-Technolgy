package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = make([][]int, 5)
	describe(i)

	var j interface{}
	i = j
	describe(i)

	var iii interface{} = "dipankar"
	t, ok := iii.(string)
	fmt.Printf("%v %v\n", t, ok)

	t1, ok1 := iii.(int)
	fmt.Printf("%v %v\n", t1, ok1)

	f := i.(float64) // panic so we need to have a boolean handlaer to avoid panic
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
