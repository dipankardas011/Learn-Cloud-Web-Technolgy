package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// Abs /** it is the method to the struct of Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) scale(mag int64) {
	v.X *= float64(mag)
	v.Y *= float64(mag)
}

func scale(v *Vertex, mag float64) {
	v.X *= mag
	v.Y *= mag
}

// Abs /** it is the method without the struct type but of type MyFloat
func (F MyFloat) Abs() float64 {
	if F < 0 {
		return float64(-F)
	} else {
		return float64(F)
	}
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//fmt.Println("Enter your name")
	//var name string
	//fmt.Scanln(&name)
	//_, err := fmt.Scanln(&name, " ")
	//if err != nil {
	//	return
	//}

	//fmt.Println("Enter your age")
	//var age uint
	//_, err1 := fmt.Scan(&age)
	//if err1 != nil {
	//	return
	//}
	//fmt.Printf("Name: <%v>, age: {%v}\n", name, age)

	v := Vertex{3, 4}
	v.scale(10)
	//scale(&v, 10.0)
	fmt.Println(v.Abs())

	f := MyFloat(-34.3)
	fmt.Println(f.Abs())

	fmt.Println(math.Sqrt2)

	v1 := Vertex{4, 5}
	fmt.Println(&v1)
	//scale(&v1, 10)
	p := &v1
	k := p.Abs()
	fmt.Println(k)
	fmt.Println(p)
	fmt.Println(AbsFunc(*p))
}
