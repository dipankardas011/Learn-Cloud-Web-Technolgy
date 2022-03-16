package main

import (
	"fmt"
	"math"
)

type Abser interface {
	abs() float64
}

type Vertex1 struct {
	x int
	y int
}

func (v *Vertex1) abs() float64 {
	return math.Sqrt(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2))
}

type Myfloat float64

func (f Myfloat) abs() float64 {
	return math.Abs(float64(f))
}

// I //////////////////////////////////////////////////////
type I interface {
	M()
	describe()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("Value <nil>")
		return
	}
	fmt.Println("Value ", t.S)
}

func (t *T) describe() {
	t.M()
	fmt.Printf("Value: %v \t Type %T\n", t, t)
}

func main() {
	var boilerPlate Abser
	boilerPlate = Myfloat(-34.0)
	fmt.Printf("%v \t %T\n", boilerPlate, boilerPlate)

	fmt.Println(boilerPlate.abs())

	boilerPlate = &Vertex1{3, 4}
	fmt.Printf("%v \t %T\n", boilerPlate, boilerPlate)

	fmt.Println(boilerPlate.abs())

	var ii I
	var ttt *T
	ii = ttt
	ii.describe()
	/**
	var ii2 I
	//	it will generate runtime error
	// as the type is no type inside the interface tuple to indicate which concrete method to call.
	ii2.describe()
	*/

}
