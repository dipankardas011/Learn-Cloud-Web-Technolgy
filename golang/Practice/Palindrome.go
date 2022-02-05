package main

import "fmt"

func isPalindrome(str string) bool {
	l := 0
	r := len(str) - 1

	for l < r {

		if str[l] != str[r] {
			return false
		}

		l++
		r--
	}
	return true
}

func main() {
	str := "Hello"
	str1 := "ABCBA"
	fmt.Println(isPalindrome(str))
	fmt.Println(isPalindrome(str1))
}
