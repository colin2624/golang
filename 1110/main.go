package main

import (
	"fmt"
)

// func supperAdd(numbers ...int) int {
// for index,  number := range numbers {
// fmt.Println(index, number)
// }

/// for i := 0; i < len(numbers); i++ {
// fmt.Println(numbers[i])
// }ㄷ

// total := 0
// for _, number := rnage numbers {
// total += numbers
// }
// return total
//  }

func cnaIDrink(age int) bool {
	if koreaAge := age + 2; koreaAge < 18 {
		return false
	} else {
		return true
	}
}

func main() {
	// supperAdd(1, 2, 3, 4, 5, 6)
	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)

	fmt.Println(cnaIDrink(16))
}
