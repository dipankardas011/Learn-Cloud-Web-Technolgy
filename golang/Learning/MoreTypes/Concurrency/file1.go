package main

import (
	"fmt"
	"time"
)

func testThreads(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		//fmt.Println("Current Time", time.Now())
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	go testThreads("ch")
	testThreads("pa")
}
