package main

import "fmt"

type student struct {
	name string
	age  int32
}

var class map[int32]student

type vertex struct {
	Lat, Long float64
}

var m = map[string]vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func mutate() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	class = make(map[int32]student)
	class[20051554] = student{
		"Dipankar Das", 21,
	}
	class[20051551] = student{
		"XYZ", 20,
	}
	fmt.Println(class)

	fmt.Println(m)

	mutate()
}
