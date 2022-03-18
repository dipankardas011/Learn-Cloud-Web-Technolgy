package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 5, 5))
	fmt.Println(m.Bounds())
	for i := m.Rect.Min.X; i < m.Rect.Max.X; i++ {
		for j := m.Rect.Min.Y; j < m.Rect.Max.Y; j++ {
			fmt.Print("(")
			fmt.Print(m.At(0, 0).RGBA())
			fmt.Print(") ")
		}
		fmt.Println()
	}
}
