package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Emp:", a)

	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:", a[4])
	fmt.Println("len:",len(a))

	b :=[5]int{1,2,3,4,5}
	fmt.Println("arr:",b)
// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	var c = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", c)

	var twoD [2][3] int
	for i := range 2 {
		for j := range 3{
			twoD[i][j]= i+j
		}
	}
	fmt.Println("2nd:",twoD)

	twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)

}