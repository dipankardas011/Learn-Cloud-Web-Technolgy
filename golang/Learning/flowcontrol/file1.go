package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

//Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
//except that Go only runs the selected case, not all the cases that follow.
//In effect, the break statement that is needed at the end of each case in those
//languages is provided automatically in Go. Another important difference is that
//Go's switch cases need not be constants, and the values involved need not be integers.
func doingSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

/**
*(Note: If you are interested in the details of the algorithm, the z² − x *
above is how far away z² is from where it needs to be (x), and the division by *
2z is the derivative of z², to scale how much we adjust z by how quickly z² is
changing. This general approach is called * Newton's method. It works well for
many functions but especially well for square root.) * z -= (z*z - x) / (2*z)
*/
func Sqrt(x float64) float64 {
	z := float64(1)
	for iter := 0; iter < 10; iter++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

/**
Variables declared inside an if short statement are also available inside any of the else blocks.
*/
func main() {
	doingSwitch()
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	fmt.Printf("Square root(3): %.4v", Sqrt(3))
}
