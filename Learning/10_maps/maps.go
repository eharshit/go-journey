//  Maps are Go’s built-in associative data type

package main

import "fmt"

func main() {
	// create map
	m:= make(map[string]string)

	// Set key/value pairs
	m["name"]="harshit"
	m["lastname"]="singh"

	// get
	fmt.Println(m["name"],m["lastname"])

	v := make(map[string]int)
	v["age"]=22
	v["price"]= 100
	// check len
	fmt.Println(len(v)) 

	// delete removes key/value pairs
	delete(v, "price")
	fmt.Println("map:", v)

	// remove all key/value pairs from a map
	clear(v)
    fmt.Println("map:", v)

	// another way to declare and initialize a new map
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)



}