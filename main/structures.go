package main

import (
	"github.com/skysoft-atm/bbl-golang/structures"
)

type pointWithAltitude struct {
	structures.Point
	altitude float32
}

type a struct {
	foo int
	b   b
}

type b struct {
	bar int
	c   c
}

type c struct {
	baz int
}

func byValue(p structures.Point) {
	p.X = 10
}

func byReference(p *structures.Point) {
	p.X = 10
}

func main() {
	/*
		A structure is a typed collection of field
	*/

	/*
		Two ways to instantiate a structure
	*/

	// First way (instance allocation)
	structureInstance := structures.Point{X: 1, Y: 2}

	// Second way (pointer allocation)
	structurePointer := new(structures.Point)
	structurePointer.X = 1
	structurePointer.Y = 2

	// Passing by value
	byValue(structureInstance)
	byValue(*structurePointer)

	// Passing by reference
	byReference(structurePointer)
	byReference(&structureInstance)

	/*
		Passing by value: immutability (copy)
		Passing by reference: performance (pointer)
	*/

	/*
		No inheritance, composition only
	*/
	p := pointWithAltitude{}
	p.X = 1
	p.Y = 2
	p.altitude = 32.4

	/*
		TODO
		Create a structure representing a customer (name, lastname, age)
		Create a function returning a customer pointer from a list of inputs
	*/
}
