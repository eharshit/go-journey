package main

	
import (
    "fmt"
    "math"
)

// const declares a constant value.
const s string = "constant"
func main() {
	fmt.Println(s)

	//const statement can also appear inside a function body.
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
