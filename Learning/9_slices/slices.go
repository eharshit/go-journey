package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil and has length 0
	var nums []string
    fmt.Println("uninit:", nums, nums == nil, len(nums) == 0)

	// use make to create a slice with non-zero length
	s := make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	
	s[0] ="a"
	s[1] ="b"
	s[2] ="c"
	fmt.Println("set:",s,"get index 2:",s[2],"len:",len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy function
	c:= make([]string,len(s))
	copy(c,s)
	fmt.Println("Copy:",c)

	// slice operator
	l := s[2:5]
    fmt.Println("sl1:", l)
	l = s[:5]
    fmt.Println("sl2:", l)
	l = s[2:]
    fmt.Println("sl3:", l)

	// declare and initialize a var in single line
	t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

	//slices
	t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
	}
	
	// 2D array in slice
	var f = [][]int{{1,2,3},{1,2,3}}
	fmt.Println("2d array in slices:",f)
	

}