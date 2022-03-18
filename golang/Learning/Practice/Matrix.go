package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func diagonal(mat [][]uint32) {
	for i, val := range mat {
		fmt.Println(i, " : ", val)
	}
}

func dis(mat [][]uint32) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}

func BinarySearch(arr []int32, left int32, right int32, target int32) int32 {
	if right < left {
		return -1
	}

	midIdx := left + (right-left)>>1
	if arr[midIdx] == target {
		return midIdx
	}
	if arr[midIdx] < target {
		return BinarySearch(arr, midIdx+1, right, target)
	}
	return BinarySearch(arr, left, midIdx-1, target)
}
func AppendByte(arr []byte, data ...byte) []byte {
	m := len(arr)
	n := m + len(data)
	if n > cap(arr) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, arr)
		arr = newSlice
	}
	arr = arr[0:n]
	copy(arr[m:n], data)
	return arr
}

func fn(x int) bool {
	return x&1 == 1
}

func filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

var digitRegexp = regexp.MustCompile("[0-9]+")

/**
This code behaves as advertised, but the returned
[]byte points into an array containing the entire file.
Since the slice references the original array, as
long as the slice is kept around the garbage
collector can’t release the array; the few useful bytes
of the file keep the entire contents in memory.
*/
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	var mat [][]uint32
	m := 4
	n := 5
	mat = make([][]uint32, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]uint32, n)
		for j := 0; j < n; j++ {
			mat[i][j] = 1
		}
	}
	dis(mat)
	diagonal(mat)

	var arr []int32
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	arr = append(arr, 7)
	arr = append(arr, 9)
	//arr := []int32{2,3,4,7,9}
	fmt.Println(arr)

	fmt.Println(BinarySearch(arr, 0, int32(len(arr)-1), 4))
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)

	fmt.Println(filter([]int{1, 2, 3, 5}, fn))

	fmt.Println("---------")
	fmt.Println(FindDigits("aaa.txt"))
}
