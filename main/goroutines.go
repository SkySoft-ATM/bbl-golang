package main

func something() {
	// Do something
}

func main() {
	something()    // Current goroutine
	go something() // Starts a new goroutine

	/*
		A goroutine is a lightweight thread (not mapped on an OS thread but a soft thread)

		The Go execution environment decides how many threads are required
		1 thread -> * goroutines

		For a CPU-bound & parallelizable operation, the optimal number of threads is ...?

		C10k problem
		Impacts: thread memory consumption, thread creation & thread switching, TLAB size in Java etc.

		Benefits: lighter, faster to start, faster to switch
	*/

	/*
		TODO
		Create a function to print a string and call it in another goroutine
	*/

	/*
		How to communicate with a goroutine?
	*/
}
