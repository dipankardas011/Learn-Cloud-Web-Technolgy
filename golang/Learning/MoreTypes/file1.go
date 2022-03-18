package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Process struct {
	pid  int16
	ppid int16
	name string
}

func main() {
	var ptr *int

	i := 42
	ptr = &i

	fmt.Println(*ptr)
	*ptr = 34
	fmt.Println(*ptr)
	fmt.Println(Vertex{3, 4})

	var myProc *Process

	myProc = new(Process)

	myProc.ppid = 1
	(*myProc).pid = 2
	myProc.name = "systemd"
	fmt.Println(myProc)

	fmt.Println(v1, p, v2, v3)

}
