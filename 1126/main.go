package main

import "fmt"

func main() {

	// a := 2
	// b := a
	// a = 10
	// fmt.Println(&a, &b)

	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}
