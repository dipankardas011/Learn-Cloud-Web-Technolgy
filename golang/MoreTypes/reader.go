package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//var b byte = 3
	//var a int = 34
	//fmt.Println(unsafe.Sizeof(b))
	//fmt.Println(unsafe.Sizeof(a))
	r := strings.NewReader("Hello World!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:] = %q\n", b[:])
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
