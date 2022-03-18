package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Millisecond)
	burst := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("💣TICK!!")
		case <-burst:
			fmt.Println("💥((( BOOM )))")
			return

		default:
			fmt.Println(" ....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
