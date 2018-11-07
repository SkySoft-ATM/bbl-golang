package main

import "fmt"

func main() {
	// Initializing a map
	m := make(map[int]string)
	m[0] = "zero"
	m[1] = "one"
	fmt.Printf("%v\n", m)

	// Other way
	m = map[int]string{
		0: "zero",
		1: "one",
	}
	fmt.Printf("%v\n", m)

	// Iteration
	for k, v := range m {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}

	// Contains
	if v, contains := m[0]; contains {
		fmt.Printf("%v\n", v)
	}

	/*
		TODO
		Enrich the customer function to cache the created customers in a map (key: name + " " + lastname)
		If the customer already exists, return the cached value
	*/
}
