package main

import "fmt"

func main() {
	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in last-in-first-out order.

	// A defer statement pushes a function call onto a list.
	// The list of saved calls is executed after the surrounding function returns.
	// Defer is commonly used to simplify functions that perform various clean-up actions.

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
