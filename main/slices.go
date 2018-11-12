package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	var slice []string

	/*
		Slicing an existing array (reference to an array)
	*/

	slice = array[0:5]
	slice = array[:]  // low:high
	slice = array[3:] // 3:high

	// A slice is a pointer to an array (plus a length and a capacity)
	array[4] = "m"
	fmt.Printf("array=%v, slice=%v\n", array, slice) // a, b, c, d, m and d, m

	slice = append(slice, "x") // Adding a new element to a slice
	fmt.Printf("%v\n", slice)  // d, m, x

	/*
		Length vs capacity
	*/

	s1 := make([]int, 0)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("%v\n", s1) // [1 2 3]

	s2 := make([]int, 0, 3)
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	fmt.Printf("%v\n", s2) // [1 2 3]

	s3 := make([]int, 3, 3)
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Printf("%v\n", s3) // [1 2 3]

	s4 := make([]int, 3, 3)
	s4 = append(s4, 1)
	s4 = append(s4, 2)
	s4 = append(s4, 3)
	fmt.Printf("%v\n", s4) // [0 0 0 1 2 3]
}
