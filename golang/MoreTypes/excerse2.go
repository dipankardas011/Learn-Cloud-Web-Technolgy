package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	list := strings.Fields(s)
	for _, v := range list {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
