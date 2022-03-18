package main

import (
	"fmt"
	"math"
)

func areaS(length int32) int32 {
	return length * length
}

func areaC(radius float32) float32 {
	return math.Pi * radius
}

func main() {
	fmt.Println("area square: ", areaS(3))
	fmt.Println("area circle: ", areaC(3.4))
}
