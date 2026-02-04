package main

import "fmt"

func main() {
// for acting like a while
	w := 5
	for w <= 9 {
    fmt.Println(w)
    w = w + 1
	}

// classic for loop 
	for i := 1; i<= 3; i++ {
		fmt.Println(i)
	}

// for loop using range on int
	for j := range 3{
		fmt.Println("range",j)
	}

// range loop skipping even numbers
	for n := range 6{
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}