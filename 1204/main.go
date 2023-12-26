package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// names := [5]string{"nice", "lynn", "dal"}
	// fmt.Println(names)

	//nico := map[string]string{"name": "nico", "age": "12"}
	//fmt.Println(nico)

	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico.name)
}
