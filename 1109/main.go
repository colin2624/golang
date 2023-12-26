package main

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
// return a * b
// }

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// func repeatMe(words ...string) {
// fmt.Println(words)
// }

func main() {
	// var name string = "nico"
	// name := "nico"
	// fmt.Println(name)
	// fmt.Println(multiply(2, 2))

	// repeatMe("inco", "lynn", "dal", "marl", "flynn")

	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)

}
